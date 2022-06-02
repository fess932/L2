package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f, err := os.Open("./develop/dev03/task.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	linuxSort(f, os.Stdout)
}

const (
	useK = iota + 1 //-k — указание колонки для сортировки
	useN            //-n — сортировать по числовому значению
	useR            //-r — сортировать в обратном порядке
	useU            //-u — не выводить повторяющиеся строки
)

type Slice []string

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func linuxSort(r io.Reader, w io.Writer, args ...int) {
	slice := make(Slice, 0, 50)

	buf := bufio.NewScanner(r)
	for buf.Scan() {
		slice = append(slice, buf.Text())
	}

	{
		sort.Strings(slice)

		for _, v := range args {
			switch v {
			case useK:
			case useR:
				slice = sort.Reverse(&slice).(Slice)
			default:
				log.Println("unknown option", v)
			}
		}

	}

	for _, v := range slice {
		io.WriteString(w, v+"\n") //nolint:errcheck
	}
}
