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
		{"", "", false},
		{"a", "a", false},
		{"abcd", "abcd", false},
		{"4a", "", true},
		{"4aaa", "", true},
		{"44", "", true},
		{"444", "", true},
		{"a2b11", "aabbbbbbbbbbb", false},
		{"a2b11c", "aabbbbbbbbbbbc", false},
		{"a2", "aa", false},
		{"a2b3c", "aabbbc", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\44`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
	}

	//t.Parallel()

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			//t.Parallel()

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
