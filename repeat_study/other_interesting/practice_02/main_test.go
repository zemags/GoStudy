package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountWordsFreq(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
		freq     int
	}{
		{
			input:    `John is son is, John second. son-of John second is, second`,
			expected: []string{"John", "is,", "second", "is"},
			freq:     4,
		},
	}

	for _, test := range tests {
		result, err := countWordsFreq(test.input, test.freq)
		require.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestCountWordsFreqInvalid(t *testing.T) {
	_, err := countWordsFreq("", 0)
	require.Error(t, err)
}
