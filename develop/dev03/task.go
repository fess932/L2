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
	// filters ops
	useU = iota + 1 //-u — не выводить повторяющиеся строки

	// sort flags
	useK //-k — указание колонки для сортировки
	useN //-n — сортировать по числовому значению
	useR //-r — сортировать в обратном порядке
)

// 1) all filter operations
// 2) all sort operations

func linuxSort(r io.Reader, w io.Writer, args ...int) {
	sort.Ints(args) // sort flags in ops order

	slice := make(sort.StringSlice, 0, 50)

	buf := bufio.NewScanner(r)
	for buf.Scan() {
		slice = append(slice, buf.Text())
	}

	var sl sort.Interface = slice
	sort.Sort(sl)

	if len(args) > 0 {
		for _, v := range args {
			switch v {
			case useU: // make uniq
				sl = Uniq(sl)
			case useK: // sort by n word in line
				// sort.Sort(byK(sl))

			case useR: // make reverse
				sl = sort.Reverse(sl)

			default:
				log.Println("unknown option", v)
			}
		}

		sort.Sort(sl)
	}

	for i := 0; i < sl.Len(); i++ {
		io.WriteString(w, slice[i]+"\n") //nolint:errcheck
	}
}

// Uniq input sorted slice, returns slice with Length() = all uniq values
func Uniq(data sort.Interface) sort.Interface {
	// calc uniq slice, set max len = last uniq index

	sort.Sort(data)

	length := data.Len()
	if length < 2 {
		return data
	}

	i, j := 0, 1

	// find the first duplicate
	for j < length && data.Less(i, j) {
		i++
		j++
	}

	// this loop is simpler after the first duplicate is found
	for ; j < length; j++ {
		if data.Less(i, j) {
			i++
			data.Swap(i, j)
		}
	}

	return &uniq{i + 1, data}
}

type uniq struct {
	lastUniq int
	sort.Interface
}

func (u uniq) Len() int {
	return u.lastUniq
}
