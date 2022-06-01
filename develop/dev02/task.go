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
