package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

const (
	// filters ops
	useU = iota + 1 //-u — не выводить повторяющиеся строки

	// sort flags
	useK //-k — указание колонки для сортировки
	useN //-n — сортировать по числовому значению
	useR //-r — сортировать в обратном порядке
)

func main() {
	f, err := os.Open("./develop/dev03/task.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	linuxSortOptions(f, os.Stdout)
}

func linuxSortOptions(r io.Reader, w io.Writer, options ...SortOption) {
	slice := make(sort.StringSlice, 0, 32)

	buf := bufio.NewScanner(r)
	for buf.Scan() {
		slice = append(slice, buf.Text())
	}

	var sl sort.Interface = slice

	for _, o := range options {
		sl = o(sl)
	}

	sort.Sort(sl)

	for i := 0; i < sl.Len(); i++ {
		io.WriteString(w, slice[i]+"\n") //nolint:errcheck
	}
}

///////////////////////////////////////////
// sort options
///////////////////////////////////////////

type SortOption func(sort.Interface) sort.Interface

func withReverse() SortOption {
	return sort.Reverse
}

////////////////////////////////////////

func withUniq() SortOption {
	return Uniq
}

// Uniq input sorted slice, returns slice with Length() == all uniq values
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

////////////////////////////////////////

func withColumnNumber(k int) SortOption {
	return func(s sort.Interface) sort.Interface {
		return &columnSorter{k - 1, s.(sort.StringSlice)}
	}
}

type columnSorter struct {
	column int
	sort.StringSlice
}

func (s columnSorter) Less(i, j int) bool {
	var str1, str2 string

	v1 := strings.SplitN(s.StringSlice[i], " ", s.column+1)
	if len(v1)-1 == s.column {
		str1 = v1[s.column]
	}

	v2 := strings.SplitN(s.StringSlice[j], " ", s.column+1)
	if len(v2)-1 == s.column {
		str2 = v2[s.column]
	}

	return str1 < str2
}

//-n — сортировать по числовому значению
func withNumber(column int) SortOption {
	return func(s sort.Interface) sort.Interface {
		return &numberSorter{column - 1, s.(sort.StringSlice)}
	}
}

type numberSorter struct {
	column int
	sort.StringSlice
}

func (s numberSorter) Less(i, j int) bool {
	var (
		buf        int
		num1, num2 int
		err        error
	)

	v1 := strings.SplitN(s.StringSlice[i], " ", s.column+1)
	if len(v1)-1 == s.column {
		buf, err = strconv.Atoi(v1[s.column])
		if err == nil {
			num1 = buf
		} else {
			log.Println(err)
		}
	}

	v2 := strings.SplitN(s.StringSlice[j], " ", s.column+1)
	if len(v2)-1 == s.column {
		buf, err = strconv.Atoi(v2[s.column])
		if err == nil {
			num2 = buf
		} else {
			log.Fatal(err)
		}
	}

	return num1 < num2
}

//-M — сортировать по названию месяца
//-b — игнорировать хвостовые пробелы
//-c — проверять отсортированы ли данные
//-h — сортировать по числовому значению с учётом суффиксов
