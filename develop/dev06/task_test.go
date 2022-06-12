package main

import (
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_parseFields(t *testing.T) {
	tcs := []struct {
		name    string
		arg     string
		want    []int
		wantErr bool
	}{
		{
			"default",
			"1,2,3",
			[]int{1, 2, 3},
			false,
		},
		{
			"one range",
			"1,2-3",
			[]int{1, 2, 3},
			false,
		},
		{
			"range from 1 to 3",
			"-3",
			[]int{1, 2, 3},
			false,
		},
	}

	for _, v := range tcs {
		t.Run(v.name, func(t *testing.T) {
			got, err := parseFields(v.arg)
			if !v.wantErr {
				require.NoError(t, err)
			}

			require.Equal(t, v.want, got)
		})
	}
}

func Test_cut(t *testing.T) {
	f, err := os.ReadFile("test.data")
	if err != nil {
		log.Fatalln(err)
	}

	tcs := []struct {
		name      string
		fields    []int
		want      string
		separated bool
	}{
		{
			"default",
			[]int{1, 3},
			`abc efg
abc 
abccdeefg 
`,
			false,
		},
		{
			"separated",
			[]int{1, 3},
			`abc efg
abc 
`,
			true,
		},
	}

	var (
		r = new(strings.Reader)
		w = new(strings.Builder)
	)

	for _, v := range tcs {
		t.Run(v.name, func(t *testing.T) {
			r.Reset(string(f))
			w.Reset()

			cut(r, w, v.fields, " ", v.separated)

			require.Equal(t, v.want, w.String())
		})
	}
}
