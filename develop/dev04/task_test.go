package main

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestAnagram(t *testing.T) {
	t.Parallel()

	st := Anagram([]string{
		"пятак", "тяпка", "пятка",
		"аааааааааабббббббббб", "ааабббаааабббббаабб",
		"листок", "слиток", "столик",
	})

	expe := map[string][]string{
		"листок": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"},
	}

	require.Equal(t, expe, st)
	log.Printf("%+v", st)
}
