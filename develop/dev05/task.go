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

	r := strings.NewReader(`
hello world
world
hello
friends
hello
`)
	log.Println(grep(r, os.Stdout, "hello", opt))
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
ищем все совпадения regexp или string.Contains
скользящие интервалы

храним последние Before+1+After строки

Если если совпадение
*/

func grep(r io.Reader, w io.Writer, pattern string, options SearchOptions) error {
	var (
		err        error
		searchFunc func(string) bool
		scanner    = bufio.NewScanner(r)
	)

	if !options.Fixed {
		searchFunc, err = searchRegexp(pattern)
		if err != nil {
			return fmt.Errorf("invalid pattern: %v", err)
		}
	} else {
		searchFunc = searchString(pattern)
	}

	var (
		n, count int
	)

	for scanner.Scan() {
		n++

		if !searchFunc(scanner.Text()) {
			continue
		}

		count++

		if options.LineNumber {
			fmt.Fprintf(w, "%d: ", n)
		}
		fmt.Fprintln(w, scanner.Text())
	}

	return nil
}

func searchRegexp(pattern string) (s func(string) bool, err error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return func(data string) bool {
		return reg.Match([]byte(data))
	}, nil
}
func searchString(pattern string) func(string) bool {
	return func(data string) bool {
		return strings.Contains(data, pattern)
	}
}

type Entry struct {
	LineNum int
	Line    string
}
