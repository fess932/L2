package main

import (
	"testing"
)

func Test_unpack(t *testing.T) {

	tcs := []struct {
		input       string
		output      string
		errExpected bool
	}{
		{"a", "a", false},
		{"abcd", "abcd", false},
		{"4a", "", true},
		{"", "", false},
		{"a2b11", "aabbbbbbbbbbb", false},
		{"a2", "aa", false},
		{"a2b3c", "aabbbc", false},
	}

	t.Parallel()

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			out, err := unpack(tc.input)
			if tc.errExpected {
				if err == nil {
					t.Errorf("expected error, got %s", out)
				}

				return
			} else if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if out != tc.output {
				t.Errorf("\nunpack(%v) = %v\nwant %v", tc.input, out, tc.output)
			}
		})
	}
}
