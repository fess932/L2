package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк) печатать только количество строк
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	opt := SearchOptions{}
	flag.BoolVar(&opt.Count, "c", false, "only show a count of matching lines")
	flag.BoolVar(&opt.IgnoreCase, "i", false, "ignore case distinctions in patterns and data")
	flag.BoolVar(&opt.Invert, "v", false, "only show non-matching lines")
	flag.BoolVar(&opt.Fixed, "F", false, "PATTERNS are strings")
	flag.BoolVar(&opt.LineNumber, "n", false, "show line numbers in front of matching lines")

	A := flag.Uint("A", 0, "print uint lines of After context")
	B := flag.Uint("B", 0, "print uint lines of Before context")
	C := flag.Uint("C", 0, "print uint lines of output context")

	flag.Parse()

	if *A > 0 {
		opt.After = *A
	}

	if *B > 0 {
		opt.Before = *B
	}

	if *C > 0 {
		opt.After = *C
		opt.Before = *C
	}

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := grep(os.Stdin, os.Stdout, flag.Arg(0), opt); err != nil {
		log.Println(err)
	}
}

type SearchOptions struct {
	After, Before uint
	Count         bool
	Invert        bool
	IgnoreCase    bool
	Fixed         bool
	LineNumber    bool
}

/*
хранить before строк
если match печатать before + match
потом читать следующие after строки и писать их
если найдено совпадение читать следующие after строки и писать их
если не найдено совпадение то начинать следюущую итерацию сканирования
*/

func grep(r io.Reader, w io.Writer, pattern string, options SearchOptions) error {
	var (
		searchFunc, err = newSearchFunc(pattern, options)

		scanner = bufio.NewScanner(r)
		before  = NewBefore(options.Before)

		e             Entry
		linenn, count int

		printFunc = func(e Entry) {
			if options.Count {
				return
			}
			if options.LineNumber {
				fmt.Fprintf(w, "%d: ", e.LineNum)
			}
			fmt.Fprintln(w, e.Text)
		}
	)

	if err != nil {
		return fmt.Errorf("newSearchFunc: %w", err)
	}

	for scanner.Scan() {
		linenn++

		e = Entry{
			LineNum: linenn,
			Text:    scanner.Text(),
		}

		if !searchFunc(e.Text) {
			before.Push(e)

			continue
		}

		count++

		before.Drain(printFunc)
		printFunc(e)

		// print next "After" lines, if line match, counter i reset to 0
		for i := uint(0); i < options.After; i++ {
			if !scanner.Scan() {
				break
			}

			linenn++

			e = Entry{
				LineNum: linenn,
				Text:    scanner.Text(),
			}

			if !searchFunc(e.Text) {
				printFunc(e)

				continue
			}

			i = 0
			count++

			printFunc(e)
		}
	}

	if options.Count {
		fmt.Fprintln(w, count)
	}

	return nil
}

type Entry struct {
	LineNum int
	Text    string
}

func newSearchFunc(pattern string, opt SearchOptions) (sf func(string) bool, err error) {
	if !opt.Fixed {
		sf, err = searchRegexp(pattern, opt)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern: %v", err)
		}
	} else {
		sf = searchString(pattern, opt)
	}

	return
}
func searchRegexp(pattern string, opt SearchOptions) (s func(string) bool, err error) {
	if opt.IgnoreCase {
		pattern = strings.ToLower(pattern) // todo: mb regexp bugged
	}

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return func(data string) bool {
		if opt.IgnoreCase {
			data = strings.ToLower(data)
		}

		return reg.Match([]byte(data)) != opt.Invert
	}, nil
}
func searchString(pattern string, opt SearchOptions) func(string) bool {
	if opt.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	return func(data string) bool {
		if opt.IgnoreCase {
			data = strings.ToLower(data)
		}

		return strings.Contains(data, pattern) != opt.Invert
	}
}

type Before struct {
	length int
	data   []Entry
}

func NewBefore(width uint) *Before {
	return &Before{
		length: 0,
		data:   make([]Entry, width),
	}
}
func (b *Before) Push(e Entry) {
	if len(b.data) == 0 {
		return
	}

	if len(b.data) == 1 {
		b.data[0] = e
		b.length = 1

		return
	}

	// если еще нет лимита, просто записываем e в конец
	if b.length < len(b.data) {
		b.data[b.length] = e
		b.length++

		return
	}

	for i := 0; i < b.length-1; i++ {
		b.data[i] = b.data[i+1]
	}

	b.data[b.length-1] = e
}

// Drain sends all entries to fn and reset length
func (b *Before) Drain(fn func(e Entry)) {
	for i := 0; i < b.length; i++ {
		fn(b.data[i])
	}

	b.length = 0
}
