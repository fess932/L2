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
	flag.Uint("A", 0, "print uint lines of leading context")
	flag.Uint("B", 0, "print uint lines of trailing context")
	flag.Uint("C", 0, "print uint lines of output context")

	flag.Bool("c", false, "only show a count of matching lines")
	flag.Bool("i", false, "ignore case distinctions in patterns and data")
	flag.Bool("v", false, "only show non-matching lines")
	flag.Bool("F", false, "PATTERNS are strings")
	flag.Bool("n", false, "show line numbers in front of matching lines")

	flag.Parse()

	r := strings.NewReader(`
hello world
world
hello
friends
hello
`)
	log.Println(grep(r, os.Stdout, "hello"))
	log.Println(grep(r, os.Stdout, "hello", WithCount()))
}

type Option struct {
}

func WithCount() Option {
	return Option{}
}

func grep(r io.Reader, w io.Writer, pattern string, options ...Option) error {
	scanner := bufio.NewScanner(r)

	var (
		n, count int
	)

	reg, err := regexp.CompilePOSIX(pattern)
	if err != nil {
		return fmt.Errorf("invalid pattern: %v", err)
	}

	for scanner.Scan() {
		n++

		line := reg.Find(scanner.Bytes())
		if line != nil {
			count++
			fmt.Fprintf(w, "%d:%s, count: %d\n", n, line, count)
		}
	}

	return nil
}

type SearchResult struct {
	Entries []Entry
}

type Entry struct {
	LineNum int
	Line    string
}
