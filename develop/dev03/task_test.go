package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const data1 = `3 three two
1 one two
2 two one
2 two one
`

const data2 = `1 one two
2 two one
2 two one
3 three two
`

const data3 = `3 three two
2 two one
2 two one
1 one two
`

func TestLinuxSort(t *testing.T) {
	tcs := []struct {
		name, in, out string
		flags         []int
	}{
		{
			"simple sort", data1, data2, nil,
		},

		{
			"reverse sort", data1, data3, []int{useR},
		},

		{
			"uniq sort", data1,
			`1 one two
2 two one
3 three two
`,
			[]int{useU},
		},

		{
			"uniq reverse sort", data1,
			`1 one two
2 two one
3 three two
`,
			[]int{useU, useR},
		},
	}

	t.Parallel()

	for _, v := range tcs {
		v := v

		t.Run(v.name, func(t *testing.T) {
			t.Parallel()

			w := &strings.Builder{}
			linuxSort(strings.NewReader(v.in), w, v.flags...)
			require.Equal(t, v.out, w.String())
		})
	}
}
