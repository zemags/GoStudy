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

func TestShortIfStatement(t *testing.T) {
	result := shortIfStatement(2.0, 2.0, 5.0)
	expectValue := 4.0
	if result != expectValue {
		t.Errorf("Values doesnt match, exptect %[1]f, but get %[2]f", expectValue, result)
	}

	result = shortIfStatement(2.0, 2.0, 3.0)
	expectValue = 3.0
	if result != expectValue {
		t.Errorf("Values doesnt match, expect %[1]f, but get %[2]f", expectValue, result)
	}
}

func TestElseIfStatement(t *testing.T) {
	result := elseIfStatement(2, 4)
	if !result {
		t.Errorf("Expect true")
	}

	result = elseIfStatement(4, 3)
	if result {
		t.Errorf("Expect false")
	}

	result = elseIfStatement(1, 1)
	if !result {
		t.Errorf("Expect true")
	}
}

func TestSwitchStatement(t *testing.T) {
	result := switchStatement(0)
	if result != "Zero" {
		t.Errorf("")
	}
	result = switchStatement(999)
	if result != "999" {
		t.Errorf("Expect 999, got %[1]s", result)
	}
	result = switchStatement(3)
	if result != "999" {
		t.Errorf("Expect 999, got %[1]s", result)
	}
}
