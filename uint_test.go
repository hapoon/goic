package goic

import (
	"fmt"
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	var sampleStruct = struct {
		name string
	}{
		"sample",
	}

	var tests = []struct {
		in    interface{}
		out   uint
		error error
	}{
		// normal
		{"1", 1, nil},
		{2, 2, nil},
		{uint(2), 2, nil},
		// abnormal
		{"a", 0, strconv.ErrSyntax},
		{sampleStruct, 0, fmt.Errorf("type error: %T", sampleStruct)},
	}
	for _, test := range tests {
		actual, err := Uint(test.in)
		if test.error != nil {
			assert.Error(t, err, "Expects error, but error not happens")
		} else {
			assert.NoError(t, err, "Expects no error, but error happens")
			assert.Equal(t, test.out, actual, "Expects %v, but actual %v", test.out, actual)
		}
	}
}
