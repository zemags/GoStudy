package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abcd", expected: "abcd"},
		{input: "3abc", expected: ""},   // error
		{input: "45", expected: ""},     // error
		{input: "aaa10b", expected: ""}, // error
		{input: "aaa0b", expected: "aab"},
		{input: "", expected: ""},
		{input: "d\n5ab", expected: "d\n\n\n\n\nabc"},
	}
	for _, test := range tests {
		result, err := unpackString(test.input)
		require.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}
