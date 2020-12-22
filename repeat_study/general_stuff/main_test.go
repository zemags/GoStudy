package main

import "testing"

func TestAppendToList(t *testing.T) {
	trueList := []int{1}
	checkList := appendToList(trueList, 1)

	waitList := []int{1, 2}

	for i := range checkList {
		if checkList[i] != waitList[i] {
			t.Errorf("Slices doesn't match")
		}
	}

}
