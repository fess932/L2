package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

}

// Anagram return
func Anagram(dict []string) map[string][]string {
	grams := make(map[ruMap][]string)

	for _, word := range dict {
		word = strings.ToLower(word)
		m := calcMap(word)

		if _, ok := grams[m]; !ok {
			grams[m] = []string{word}

			continue
		}

		grams[m] = append(grams[m], word)
	}

	total := make(map[string][]string, len(grams))

	i := 0

	for _, v := range grams {
		if len(v) < 2 {
			continue
		}

		vals := v
		//vals := v[1:] // пропускаем первое слово из множества если надо
		sort.Strings(vals)
		total[v[0]] = vals

		i++
	}

	return total
}

// o(n^2)
func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	for _, c := range str1 {
		if !strings.Contains(str2, string(c)) {
			return false
		}
	}

	return true
}

type ruMap [34]int

// o(2n)
func isAnagram2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	return calcMap(str1) == calcMap(str2)
}

func calcMap(str string) ruMap {
	var dict ruMap

	for _, c := range str {
		dict[c-1072]++
	}

	return dict
}
