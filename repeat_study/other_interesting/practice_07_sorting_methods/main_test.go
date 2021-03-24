package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	expected := []int{0, 2, 2, 2, 3, 4, 5, 5, 8, 9}
	actual := countingSort([]int{8, 4, 0, 3, 5, 2, 2, 9, 2, 5})
	assert.Equal(t, expected, actual)
}
