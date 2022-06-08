package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CD(t *testing.T) {
	tcs := []struct {
		name   string
		path   string
		result string
	}{
		{
			name:   "user home",
			path:   "",
			result: "/home/restar",
		},
		{
			name:   "home",
			path:   "/home",
			result: "/home",
		},
		{
			name:   "root",
			path:   "..",
			result: "/",
		},
	}

	for _, v := range tcs {
		t.Run(v.name, func(t *testing.T) {
			cd(v.path)
			dir, err := os.Getwd()
			require.NoError(t, err)
			require.Equal(t, v.result, dir)
		})
	}
}
