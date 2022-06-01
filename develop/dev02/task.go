package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	log.Println(unpackV2("a1b2"))
}

const (
	number = iota + 1
	letter
	empty
)

type token struct {
	t     int
	value interface{}
}

func (t token) string() string {
	log.Println(t.value)
	return string(t.value.(rune))
}

func (t token) int() int {
	return t.value.(int)
}

type tokenizer struct {
	idx     int
	letters []rune
	current token
	next    token
}

func (t *tokenizer) readNext() {
	switch {
	case t.idx >= len(t.letters):
		t.next.t = empty

	case t.letters[t.idx] == '\\':
		t.idx++
		if t.idx >= len(t.letters) {
			t.next.t = empty

			return
		}

		t.next = token{t: letter, value: t.letters[t.idx]}

	case isNumber(t.letters[t.idx]):
		t.next = token{t: number, value: t.getCount()}

	default:
		t.next = token{t: letter, value: t.letters[t.idx]}
	}

	t.idx++
}

func (t *tokenizer) token() (next bool) {
	if t.idx > len(t.letters) {
		return false
	}

	if t.idx == 0 {
		t.readNext()
	}

	t.current = t.next
	t.readNext()

	return t.idx < len(t.letters)+1
}

func (t *tokenizer) getCount() (count int) {
	var num string

	for ; t.idx < len(t.letters); t.idx++ {
		if !isNumber(t.letters[t.idx]) {
			break
		}

		num += string(t.letters[t.idx])
	}
	t.idx--

	// must get number
	count, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (t *tokenizer) check() error {
	if len(t.letters) == 0 {
		return ErrEmptyString
	}

	if isNumber(t.letters[0]) {
		return ErrInvalidString
	}

	return nil
}

var ErrInvalidString = fmt.Errorf("invalid string")
var ErrEmptyString = fmt.Errorf("empty string")

func isNumber(str rune) bool {
	return strings.ContainsRune("1234567890", str)
}

func unpackV2(pkd string) (string, error) {
	tkz := &tokenizer{
		idx:     0,
		letters: []rune(pkd),
	}

	if err := tkz.check(); err != nil {
		if errors.Is(err, ErrEmptyString) {
			return "", nil
		}

		return "", err
	}

	var (
		total strings.Builder
	)

	for tkz.token() {
		switch {
		case tkz.next.t == number:
			if tkz.current.t == letter {
				total.WriteString(
					strings.Repeat(tkz.current.string(), tkz.next.int()),
				)
			}

		case tkz.next.t == letter:
			if tkz.current.t == letter {
				total.WriteString(tkz.current.string())
			}
		default:
			log.Println("UNKNOWN", tkz.current, tkz.next)
		}
	}

	if tkz.current.t == letter {
		total.WriteString(tkz.current.string())
	}

	return total.String(), nil
}
