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

type ttype int

const (
	num ttype = iota + 1
	str
	esc
)

type token struct { // если это число то они идут подряд
	t     ttype
	value string
}

func ntoken(s string) (t token) {
	if strings.Contains("1234567890", s) {
		t.t = num
		return t
	}

	t.t = str
	t.value = s
	return t
}

type unpacker struct {
	cursor int

	curr token
	next token

	nums    string
	letters []string
	strings.Builder
}

func (u *unpacker) Next() bool {
	if u.cursor > len(u.letters)-2 {
		return false
	}

	if u.cursor == 0 {
		u.curr = ntoken(u.letters[u.cursor])
		u.cursor++
		u.next = ntoken(u.letters[u.cursor])
	} else {
		u.curr = u.next
		u.cursor++
		u.next = ntoken(u.letters[u.cursor])
	}

	return true
}

func (u *unpacker) writeNTimes(s, n string) {
	count, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}

	u.WriteString(strings.Repeat(s, count))
}

func unpack(pkd string) (string, error) {
	if pkd == "" {
		return "", nil
	}

	if len(pkd) < 2 {
		if strings.Contains("1234567890", pkd) {
			return "", fmt.Errorf("invalid string")
		}

		return pkd, nil
	}

	letters := strings.Split(pkd, "")

	log.Println("first", letters[0])

	if strings.Contains("1234567890", letters[0]) {
		return "", fmt.Errorf("invalid string")
	}

	total := ""
	letter := ""
	for _, l := range letters {
		if !strings.Contains("1234567890", l) {
			letter = l
			continue
		}

		count, err := strconv.Atoi(l)
		if err != nil {
			return "", fmt.Errorf("invalid string")
		}

		total += strings.Repeat(letter, count)
	}

	return total, nil

	//u := unpacker{
	//	letters: strings.Split(s, ""),
	//}
	//
	//var (
	//	strNum, letter string
	//)
	//
	//for u.Next() {
	//	switch u.curr.t {
	//	case num:
	//		if letter == "" {
	//			return "", fmt.Errorf("некорректная строка")
	//		}
	//
	//		if u.next.t != num {
	//			u.writeNTimes(letter, strNum)
	//			strNum = ""
	//			letter = ""
	//		}
	//	case str:
	//		letter = u.curr.value
	//	default:
	//		log.Fatal("unknown token type")
	//	}
	//}
	//
	//return u.String(), nil
	//
	//
	//
	//
	//
	//
	////a2b33
	////
	////1 итерация
	////
	////curr = a
	////next = 2
	////
	////letter = curr
	////strNum = next
	////
	////2 итерация
	////curr = 2
	////next = b
	////
	////write(letter * strNum)
	////letter = next
	////strNum = ""
	////
	////3 итерация
	////curr = b
	////next = 3
	////
	////strNum += next
	////
	////4 итерация
	////curr = 3
	////next = 3
	////
	////strNum += next
	////
	////5 итерация
	////LAST
	////
	////write(letter * strNum)
	//
	//
	//
	//isNumber := func(s string) bool {
	//	return strings.Contains(nums, s)
	//}
	//
	//var next string
	//for i, curr := range letters {
	//	if i == len(letters)-1 { // последнее значение сразу записываем
	//		if isNumber(curr) {
	//			strNum += curr
	//		}
	//
	//		writeNTimes(letter, strNum)
	//
	//		break
	//	}
	//
	//	next = letters[i+1]
	//
	//	if !isNumber(curr) {
	//		if !isNumber(next){
	//
	//		}
	//	}
	//
	//	if !isNumber(next) {
	//		b.WriteString(curr)
	//	}
	//
	//
	//	current letter | number
	//	next    letter | number | ""
	//
	//	если current == number && next != number
	//	тогда letter надо записать с текущим номером
	//
	//	if !isNumber(next) {
	//
	//	}
	//
	//	writeNTimes(letter, strNum)
	//
	//	if isNumber(curr) {
	//		if letter != "" {
	//			writeNTimes(letter, strNum)
	//		}
	//	}
	//
	//	next = letters[i+1]
	//	if isNumber(next) {
	//		continue
	//	}
	//
	//	if i == len(letters)-1 {
	//		// 3 если текущий символ последний то записываем
	//		if !strings.Contains(nums, curr) {
	//			b.WriteString(curr)
	//
	//			break
	//		}
	//
	//		strNum += curr
	//
	//		count, err := strconv.Atoi(strNum)
	//		if err != nil {
	//			return "", fmt.Errorf("invalid string: %w", err)
	//		}
	//
	//		b.WriteString(strings.Repeat(letter, count))
	//
	//		break
	//	}
	//
	//	// 1 если текущая буква число и есть предыдущая буква
	//	// тогда добавляем число в числовую строку
	//	if strings.Contains(nums, curr) && letter != "" {
	//		strNum += curr
	//
	//		// 3 если текущая буква последняя то записываем её
	//		continue
	//	}
	//
	//	// 2 если есть предыдущая буква то записываем предыдущую строку * количество раз
	//	// в текстовый буфер количество строк затем очищаем число и строку
	//	if letter != "" {
	//		count, err := strconv.Atoi(strNum)
	//		if err != nil {
	//			return "", fmt.Errorf("invalid string: %w", err)
	//		}
	//
	//		b.WriteString(strings.Repeat(letter, count))
	//
	//		strNum = ""
	//	}
	//
	//	letter = curr
	//
	//}
	//
	//return b.String(), nil
}
