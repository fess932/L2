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

func linuxSort(r io.Reader, w io.Writer, args ...int) {
	slice := make(sort.StringSlice, 0, 50)

	buf := bufio.NewScanner(r)
	for buf.Scan() {
		slice = append(slice, buf.Text())
	}

	{
		var sl sort.Interface = slice

		if len(args) == 0 {
			sort.Strings(slice)
		}

		for _, v := range args {
			switch v {
			case useK: // sort by n word in line
				//sort.Sort(byK(sl))
			case useR: // make reverse
				sl = sort.Reverse(sl)
			case useU: // make uniq
				sl = Uniq(sl)
				log.Println(sl)
			default:
				log.Println("unknown option", v)
			}
		}

		sort.Sort(sl)
	}

	for _, v := range slice {
		io.WriteString(w, v+"\n") //nolint:errcheck
	}
}

func Uniq(data sort.Interface) sort.Interface {
	// calc uniq slice, set max len = last uniq index

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
		if !(data.Less(i, j) && data.Less(j, i)) {
			i++
			data.Swap(i, j)
		}
	}

	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	return &uniq{i + 1, data}
}

type uniq struct {
	lastUniq int
	sort.Interface
}

func (u uniq) Len() int {
	return u.lastUniq
}
