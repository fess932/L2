package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func Test_grep(t *testing.T) {
	file, err := os.ReadFile("./test.data")
	if err != nil {
		panic(err)
	}

	tcs := []struct {
		name             string
		in, pattern, out string
		opt              SearchOptions
	}{
		{
			name:    "no options",
			in:      string(file),
			pattern: "fr.ends",
			out: `friends
`,
		},
		{
			name:    "fixed bad pattern",
			in:      string(file),
			opt:     SearchOptions{Fixed: true},
			pattern: "fr.ends",
			out:     ``,
		},
		{
			name:    "fixed ok",
			in:      string(file),
			opt:     SearchOptions{Fixed: true},
			pattern: "friend",
			out: `friends
`,
		},
		{
			name:    "Ignore case regexp",
			in:      string(file),
			pattern: "Fr.ends",
			opt:     SearchOptions{IgnoreCase: true},
			out: `friends
`,
		},
		{
			name:    "Ignore case string",
			in:      string(file),
			pattern: "Friends",
			opt:     SearchOptions{IgnoreCase: true, Fixed: true},
			out: `friends
`,
		},
		{
			name:    "regexp line number",
			in:      string(file),
			pattern: ".at",
			opt:     SearchOptions{LineNumber: true},
			out: `1: cats
7: cats
11: cats
`,
		},
		{
			name:    "FIXED regexp line number",
			in:      string(file),
			pattern: "\n*at",
			opt:     SearchOptions{LineNumber: true, Fixed: true},
			out:     ``,
		},
		{
			name:    "FIXED string number",
			in:      string(file),
			pattern: "cat",
			opt:     SearchOptions{LineNumber: true, Fixed: true},
			out: `1: cats
7: cats
11: cats
`,
		},
		{
			name:    "context 2",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{Before: 2, After: 2},
			out: `cats
hello
world
words
words
cats
words
words
wtf
cats
`,
		},
		{
			name:    "before 1",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{Before: 1},
			out: `cats
words
cats
wtf
cats
`,
		},
		{
			name:    "after 1",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{After: 1},
			out: `cats
hello
cats
words
cats
`,
		},
		{
			name:    "after 1 with number",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{After: 1, LineNumber: true},
			out: `1: cats
2: hello
7: cats
8: words
11: cats
`,
		},
		{
			name:    "invert after 1",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{After: 1, Invert: true},
			out: `hello
world
friends
words
words
cats
words
words
wtf
cats
`,
		},

		{
			name:    "count lines",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{Count: true},
			out: `3
`,
		},

		{
			name:    "invert count lines",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{Count: true, Invert: true},
			out: `8
`,
		},
		{
			name:    "invert context 1 count lines",
			in:      string(file),
			pattern: "cats",
			opt:     SearchOptions{Count: true, Before: 1, After: 1, Invert: true},
			out: `8
`,
		},
	}

	var (
		in  = new(strings.Reader)
		out = new(strings.Builder)
	)

	for _, v := range tcs {
		v := v
		t.Run(v.name, func(t *testing.T) {
			in.Reset(v.in)
			out.Reset()

			if err = grep(in, out, v.pattern, v.opt); err != nil {
				t.Errorf("grep(%q, %q) error: %v", v.in, v.pattern, err)
			}

			require.Equal(t, v.out, out.String())
		})
	}
}
