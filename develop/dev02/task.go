package main

import (
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

func unpack(s string) (string, error) {
	const nums = "1234567890"

	var (
		b           strings.Builder
		num, letter string
		letters     = strings.Split(s, "")
	)

	if len(letters) < 2 {
		return s, nil
	}

	for _, v := range letters {
		if strings.Contains(nums, v) {
			num += v

			continue
		}

		if num == "" {
			b.WriteString(letter)

			continue
		}

		count, err := strconv.Atoi(num)
		if err != nil {
			return "", fmt.Errorf("invalid string: %w", err)
		}

		b.WriteString(strings.Repeat(letter, count))

		letter = v
	}

	return b.String(), nil
}
