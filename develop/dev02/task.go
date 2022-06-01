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
	log.Println(unpack("a1b2"))
}

const (
	number = iota + 1
	letter
)

type token struct {
	t     int
	value interface{}
}

type tokenizer struct {
	idx     int
	letters []rune
	current token
	next 	token
}

func (t *tokenizer) token() (next bool) {

	// recive current token
	switch {
	case t.letters[t.idx] == '\\':
		t.idx++
		t.current = token{t: letter, value: t.letters[t.idx]}

	case isNumber(t.letters[t.idx]):
		t.current = token{t: number, value: t.getCount()}
		t.idx++

	default:
		t.current = token{t: letter, value: t.letters[t.idx]}
		t.idx++
	}

	// recive next token
	switch {
	case t.letters[t.idx] == '\\':
		t.idx++
		t.next = token{t: letter, value: t.letters[t.idx]}

	case isNumber(t.letters[t.idx]):
		t.next = token{t: number, value: t.getCount()}

	default:
		t.next = token{t: letter, value: t.letters[t.idx]}
		t.idx++
	}

	return t.idx < len(t.letters)
}

func (t *tokenizer) getCount() (count int) {
	var num string

	for i := idx; i <= len(letters); i++ {
		end = i

		if i == len(letters) { // конец строки
			break
		}

		if !isNumber(letters[i]) {
			break
		}

		num += string(letters[i])
	}

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

//func (t *tokenizer) nextToken() token {
//	if t.escape {
//		t.escape = false
//		return token{t: letter, value: t.letters[i]}
//	}
//
//	if isNumber(t.letters[i]) {
//		return token{t: number, value: t.letters[i]}
//	}
//
//	return token{t: letter, value: t.letters[i]}
//}

var ErrInvalidString = fmt.Errorf("invalid string")
var ErrEmptyString = fmt.Errorf("empty string")

func isNumber(str rune) bool {
	return strings.ContainsRune("1234567890", str)
}

func check(str []rune) error {
	if len(str) == 0 {
		return ErrEmptyString
	}

	if isNumber(str[0]) {
		return ErrInvalidString
	}

	return nil
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
		total          strings.Builder
		count, nextIdx int
	)

	var ltr := ""

	for tkz.token() {
		switch tkz.current.t {
		case number: // число на которое надо умножить letter
		case letter: // символ который надо повторить
			ltr *
		}
	}

	//escape := false

	for i := 0; i < len(letters); i++ {
		if i == len(letters)-1 {
			total.WriteRune(letters[i])

			break
		}

		if !isNumber(letters[i+1]) {
			total.WriteRune(letters[i])

			continue
		}

		//if letters[i] == '\\' {
		//	escape = true
		//	// escape sequence
		//	// next symbol always is string
		//}

		count, nextIdx = getCount(i + 1)

		total.WriteString(strings.Repeat(string(letters[i]), count))

		if nextIdx > len(letters)-1 {
			// последний элемент, заканчивается тут
			// это означает что больше символов нет
			break
		} else {
			log.Println("next letter: ", letters[nextIdx])
		}

		i = nextIdx - 1
	}

	return total.String(), nil
}

func unpack(pkd string) (string, error) {
	letters := []rune(pkd)

	if err := check(letters); err != nil {
		if errors.Is(err, ErrEmptyString) {
			return "", nil
		}

		return "", err
	}

	// returns count of repeated symbols and next symbol after count
	getCount := func(start int) (count int, end int) {
		var num string

		for i := start; i <= len(letters); i++ {
			end = i

			if i == len(letters) { // конец строки
				break
			}

			if !isNumber(letters[i]) {
				break
			}

			num += string(letters[i])
		}

		// must get number
		count, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	var (
		total          strings.Builder
		count, nextIdx int
	)

	//escape := false

	for i := 0; i < len(letters); i++ {
		if i == len(letters)-1 {
			total.WriteRune(letters[i])

			break
		}

		if !isNumber(letters[i+1]) {
			total.WriteRune(letters[i])

			continue
		}

		//if letters[i] == '\\' {
		//	escape = true
		//	// escape sequence
		//	// next symbol always is string
		//}

		count, nextIdx = getCount(i + 1)

		total.WriteString(strings.Repeat(string(letters[i]), count))

		if nextIdx > len(letters)-1 {
			// последний элемент, заканчивается тут
			// это означает что больше символов нет
			break
		} else {
			log.Println("next letter: ", letters[nextIdx])
		}

		i = nextIdx - 1
	}

	return total.String(), nil
}
