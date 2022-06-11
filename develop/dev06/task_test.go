package main

import (
	"github.com/stretchr/testify/require"
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
			"1-3",
			[]int{1, 2, 3},
			false,
		},
		{
			"range from 1 to 3",
			"-3",
			[]int{1, 2, 3},
			false,
		},
		{
			"range from 1 to end",
			"1-",
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
