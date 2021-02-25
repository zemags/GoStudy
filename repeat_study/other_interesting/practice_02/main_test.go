package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWordsFreq(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
		freq     int
	}{
		{input: `John is the son of John second.
         Second son of John second is William second`, expected: []string{"second", "John", "son", "is"},
			freq: 4,
		},
	}

	for _, test := range tests {
		result := countWordsFreq(test.input, test.freq)
		assert.Equal(t, test.expected, result)
	}
}
