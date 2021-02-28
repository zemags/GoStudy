package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyTrue(t *testing.T) {
	emptyTestList := SingleLinkedStruct{}
	result := emptyTestList.IsEmpty()
	assert.Equal(t, true, result)
}

func TestIsEmptyFalse(t *testing.T) {
	notEmptyTestList := SingleLinkedStruct{}
	notEmptyTestList.value = 1
	result := notEmptyTestList.IsEmpty()
	assert.Equal(t, false, result)
}

func TestAddForward(t *testing.T) {
	testList := SingleLinkedStruct{}
	testList.AddForward(1)
}
