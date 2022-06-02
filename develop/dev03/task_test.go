package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const data1 = `3 three two 10
1 one two 2
2 two one
2 two one
`

func TestSortSlice(t *testing.T) {
	tcs := []struct {
		name, in, out string
		flags         []SortOption
	}{
		{
			"simple sort",
			`3 three two 10
1 one two 2
2 two one
2 two one
`,
			`1 one two 2
2 two one
2 two one
3 three two 10
`, nil,
		},
		{
			"reverse sort",
			`3 three two 10
1 one two 2
2 two one
2 two one`,
			`3 three two 10
2 two one
2 two one
1 one two 2
`, []SortOption{withReverse()},
		},
		{
			"uniq sort", data1,
			`1 one two 2
2 two one
3 three two 10
`,
			[]SortOption{withUniq()},
		},
		{
			"uniq reverse sort", data1,
			`3 three two 10
2 two one
1 one two 2
`,
			[]SortOption{withUniq(), withReverse()},
		},
		{
			"sort k column number", data1,
			`2 two one
2 two one
3 three two 10
1 one two 2
`,
			[]SortOption{withColumnNumber(4)},
		},
		{
			"sort k column number rever", data1,
			`1 one two 2
3 three two 10
2 two one
`,
			[]SortOption{withColumnNumber(4), withReverse(), withUniq()},
		},
		{
			"sort by number value and column 4",
			`3 three two 10
1 one two 2s s
2 two one
2 two one`,
			`2 two one
2 two one
1 one two 2s s
3 three two 10
`,
			[]SortOption{withNumber(4)},
		},
	}

	t.Parallel()

	for _, v := range tcs {
		v := v

		t.Run(v.name, func(t *testing.T) {
			t.Parallel()

			w := &strings.Builder{}
			Sort(strings.NewReader(v.in), w, v.flags...)
			require.Equal(t, v.out, w.String())
		})
	}
}
