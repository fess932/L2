package main

import (
	"testing"
)

func Fuzz_unpack2(f *testing.F) {
	f.Add(5, "hello")
	f.Fuzz(func(t *testing.T, i int, s string) {
		out, err := unpackV2(s)
		if err != nil && out != "" {
			t.Errorf("%q, %v", out, err)
		}
	})
}

func Test_unpack2(t *testing.T) {

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
		{`qwe\44`, "qwe4444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`qwe\\5\\`, `qwe\\\\\\`, false},
		{`qwe\\5\`, `qwe\\\\\`, false},
		{`f1351a51c9fa1633a2757bbc648d0fccd38523dd2ce28c803b3eff822bf99e80`, "", false},
	}

	//t.Parallel()

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			//t.Parallel()

			out, err := unpackV2(tc.input)
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
