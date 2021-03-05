package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetLenValid(t *testing.T) {
	testList := &Node{}
	testList.AddBackword(1)
	_, err := testList.GetLen()
	require.NoError(t, err)
}

func TestGetLenInvalid(t *testing.T) {
	testList := &Node{}
	_, err := testList.GetLen()
	require.Error(t, err)
}

func TestGetFirstValid(t *testing.T) {
	testList := &Node{value: 1}
	_, err := testList.GetFirst()
	require.NoError(t, err)
}

func TestGetFirstInvalid(t *testing.T) {
	testList := &Node{}
	_, err := testList.GetFirst()
	require.Error(t, err)
}

func TestGetAFirstValid(t *testing.T) {
	expected := []int{2, 3}
	testList := &Node{}
	for _, i := range []int{1, 2, 3} {
		testList.AddBackword(i)
	}
	testList, err := testList.GetAFirst()
	require.NoError(t, err)
	actual, _ := DisplayList(testList)
	assert.Equal(t, expected, actual)
}

func TestGetAFirstInvalid(t *testing.T) {
	testList := &Node{}
	_, err := testList.GetAFirst()
	require.Error(t, err)
}

func TestDisplayListValid(t *testing.T) {
	expected := []int{1, 2, 3}
	testList := &Node{}
	for _, i := range []int{1, 2, 3} {
		testList.AddBackword(i)
	}
	actual, err := DisplayList(testList)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestTestDisplayListInvalid(t *testing.T) {
	testList := &Node{}
	_, err := DisplayList(testList)
	require.Error(t, err)
}

func TestAddForward(t *testing.T) {
	expected := []int{4, 1, 2, 3}
	testList := &Node{}
	for _, i := range []int{1, 2, 3} {
		testList.AddBackword(i)
	}
	testList.AddForward(4)
	actual, _ := DisplayList(testList)
	assert.Equal(t, expected, actual)
}

func TestAddInto(t *testing.T) {
	expected := []int{1, 2, 5, 3, 4}
	testList := &Node{}
	for _, i := range []int{1, 2, 3, 4} {
		testList.AddBackword(i)
	}
	err := testList.AddInto(5, 2)
	require.NoError(t, err)
	actual, _ := DisplayList(testList)
	assert.Equal(t, expected, actual)
}

func TestAddIntoInvalidIdx(t *testing.T) {
	testList := &Node{}
	err := testList.AddInto(1, 1)
	require.Error(t, err)
}

func TestAddIntoIdxOutOfRange(t *testing.T) {
	testList := &Node{}
	testList.AddBackword(1)
	err := testList.AddInto(2, 3)
	require.Error(t, err)
}

func TestRemove(t *testing.T) {
	expected := []int{2, 3, 4}
	testList := &Node{}
	for _, i := range []int{1, 2, 3, 4} {
		testList.AddBackword(i)
	}
	err := testList.Remove(0)
	require.NoError(t, err)
	actual, _ := DisplayList(testList)
	assert.Equal(t, expected, actual)

	expected = []int{2, 4}
	err = testList.Remove(1)
	require.NoError(t, err)
	actual, _ = DisplayList(testList)
	assert.Equal(t, expected, actual)
}

func TestRemoveEmpty(t *testing.T) {
	testList := &Node{}
	err := testList.Remove(1)
	require.Error(t, err)
}

func TestRemoveIndexOutOfRange(t *testing.T) {
	testList := &Node{}
	testList.AddBackword(1)
	err := testList.Remove(2)
	require.Error(t, err)
	err = testList.Remove(-1)
	require.Error(t, err)
}
