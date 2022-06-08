package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CD(t *testing.T) {
	os.Setenv("HOME", "/home/fess932")

	tcs := []struct {
		name   string
		path   string
		result string
	}{
		{
			name:   "home",
			path:   "",
			result: "/home/fess932",
		},
	}

	for _, v := range tcs {

		t.Run(v.name, func(t *testing.T) {
			cd(v.path)
			require.Equal(t, v.result, os.Getenv("PWD"))
		})
	}
}
