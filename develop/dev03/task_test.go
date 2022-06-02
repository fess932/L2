package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const data1 = `3 three two
1 one two
2 two one
`

const data2 = `1 one two
2 two one
3 three two
`

func TestLinuxSort(t *testing.T) {
	w := &strings.Builder{}

	linuxSort(
		strings.NewReader(data1),
		w,
	)

	require.Equal(t, w.String(), data2)
}
