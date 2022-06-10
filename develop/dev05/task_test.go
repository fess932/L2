package main

import (
	"strings"
	"testing"
)

func Test_grep(t *testing.T) {
	// Проверка по умолчанию
	opt := SearchOptions{}
	r := strings.NewReader(`
hello world
world
hello
friends
hello
`)
	if grep(r, "hello", opt) != `hello
hello
hello
` {
		t.Error("grep default")
	}
}
