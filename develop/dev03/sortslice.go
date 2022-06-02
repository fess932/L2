package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Sort(r io.Reader, w io.Writer, options ...SortOption) {
	slice := make(SortSlice, 0, 32)

	buf := bufio.NewScanner(r)
	for buf.Scan() {
		slice = append(slice, strings.Split(buf.Text(), " "))
	}

	var sl sort.Interface = slice

	for _, o := range options {
		sl = o(sl)
	}

	sort.Sort(sl)

	for i := 0; i < sl.Len(); i++ {
		io.WriteString(w, slice[i].String()+"\n") //nolint:errcheck
	}
}

type SortRow []string

func (s SortRow) String() string {
	return strings.Join(s, " ")
}

// SortSlice attaches the methods of Interface to []SortRow, sorting in increasing order.
type SortSlice []SortRow

func (x SortSlice) Len() int           { return len(x) }
func (x SortSlice) Less(i, j int) bool { return x[i].String() < x[j].String() }
func (x SortSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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
// calc uniq slice, set max len = last uniq index
func Uniq(data sort.Interface) sort.Interface {
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
		return &columnSorter{k - 1, s.(SortSlice)}
	}
}

type columnSorter struct {
	column int
	SortSlice
}

func (s columnSorter) Less(i, j int) bool {
	var str1, str2 string

	if len(s.SortSlice[i]) > s.column {
		str1 = s.SortSlice[i][s.column]
	}

	if len(s.SortSlice[j]) > s.column {
		str2 = s.SortSlice[j][s.column]
	}

	return str1 < str2
}

////////////////////////////////////////

//-n — сортировать по числовому значению
func withNumber(column int) SortOption {
	return func(s sort.Interface) sort.Interface {
		return &numberSorter{column - 1, s.(SortSlice)}
	}
}

type numberSorter struct {
	column int
	SortSlice
}

func (s numberSorter) Less(i, j int) bool {
	var num1, num2 int

	if len(s.SortSlice[i]) > s.column {
		num1 = parseInt(s.SortSlice[i][s.column])
	}

	if len(s.SortSlice[j]) > s.column {
		num2 = parseInt(s.SortSlice[j][s.column])
	}

	return num1 < num2
}

// parseInt parse bad string with first number chars to int
// "9asd" > "9"
func parseInt(str string) int {
	total := ""
	for _, c := range str {
		// todo: else "-" for negative numbers
		if !strings.ContainsAny(string(c), "0123456789") {
			break
		}

		total += string(c)
	}

	num, _ := strconv.Atoi(total)

	return num
}

//-M — сортировать по названию месяца
//-b — игнорировать хвостовые пробелы
//-c — проверять отсортированы ли данные
//-h — сортировать по числовому значению с учётом суффиксов
