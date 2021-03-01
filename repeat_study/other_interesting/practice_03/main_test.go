package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyTrue(t *testing.T) {
	emptyTestList := Node{}
	result := emptyTestList.IsEmpty()
	assert.Equal(t, true, result)
}

func TestIsEmptyFalse(t *testing.T) {
	notEmptyTestList := Node{}
	notEmptyTestList.value = 1
	result := notEmptyTestList.IsEmpty()
	assert.Equal(t, false, result)
}

func TestAddBackword(t *testing.T) {
	testList := &Node{}
	testList.AddBackword(1)
	testList.AddBackword(2)
	assert.Equal(t, 1, testList.value)
	testSecond := testList.next
	assert.Equal(t, 2, testSecond.value)
}
