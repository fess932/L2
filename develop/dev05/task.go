package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	r := strings.NewReader(`
hello world
world
hello
friends
hello
`)
	grep(r, os.Stdout, "hello")
}

type Option struct {
}

func grep(r io.Reader, w io.Writer, pattern string, options ...Option) {
	scanner := bufio.NewScanner(r)

	var n int
	for scanner.Scan() {
		n++

		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Fprintf(w, "%d:%s\n", n, line)
		}
	}
}
