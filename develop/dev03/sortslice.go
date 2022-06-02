package main

import (
	"bufio"
	"io"
	"sort"
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
