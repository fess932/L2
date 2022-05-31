package main

import (
	"testing"
)

func Test_unpack(t *testing.T) {

	tcs := []struct {
		input  string
		output string
	}{
		{"a", "a"},
		{"a2b3c", "aabbbc"},
		{"a2b3c4d", "aabbbccccd"},
	}

	t.Parallel()

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			out, err := unpack(tc.input)
			if err != nil {
				t.Errorf("\nunpack(%s)\nerror: %v", tc.input, err)

				return
			}

			if out != tc.output {
				t.Errorf("\nunpack(%v) = %v\nwant %v", tc.input, out, tc.output)
			}
		})
	}
}
