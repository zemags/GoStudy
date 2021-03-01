package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyTrue(t *testing.T) {
	emptyTestList := Node{}
	actual := emptyTestList.IsEmpty()
	assert.Equal(t, true, actual)
}

func TestIsEmptyFalse(t *testing.T) {
	notEmptyTestList := Node{}
	notEmptyTestList.value = 1
	actual := notEmptyTestList.IsEmpty()
	assert.Equal(t, false, actual)
}

func TestAddBackword(t *testing.T) {
	testList := &Node{}
	testList.AddBackword(1)
	testList.AddBackword(2)
	assert.Equal(t, 1, testList.value)
	testSecond := testList.next
	assert.Equal(t, 2, testSecond.value)
}

func TestGetLen(t *testing.T) {
	testList := &Node{}
	actual := testList.GetLen()
	assert.Equal(t, 0, actual)
}
