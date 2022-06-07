package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunc(t *testing.T) {
	tcs := []struct {
		url  string
		walk bool
	}{
		{"https://go.dev/", false},
		{"https://dev.go.dev", false},
	}

	t.Parallel()
	for i, tc := range tcs {
		i, tc := i, tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.walk, toVisit(tc.url))
		})
	}
}
