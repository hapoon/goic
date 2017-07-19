package goic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	var sampleStruct = struct {
		name string
	}{
		"sample",
	}

	sampleArray := []int{1, -2, 3, 4}

	var tests = []struct {
		in  interface{}
		out string
	}{
		{1, "1"},
		{-2, "-2"},
		{0.3, "0.3"},
		{sampleStruct, "{sample}"},
		{sampleArray, "[1 -2 3 4]"},
	}

	for _, test := range tests {
		actual := String(test.in)
		assert.Equal(t, test.out, actual, "Expects %v, but actual %v", test.out, actual)
	}
}
