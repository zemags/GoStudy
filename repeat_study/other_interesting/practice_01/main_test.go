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
		{input: "aaa0b", expected: "aab"},
		{input: "", expected: ""},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}
	for _, test := range tests {
		result, err := unpackString(test.input)
		require.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestUnpackInvalidString(t *testing.T) {
	tests := []string{"65", "1abc", "abc98nn"}
	for _, test := range tests {
		_, err := unpackString(test)
		require.Error(t, err)
	}
}
