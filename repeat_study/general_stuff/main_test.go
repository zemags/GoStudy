package main

import (
	"testing"
)

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
	if result != "Default value" {
		t.Errorf("Expect 999, got %[1]s", result)
	}
}

func TestFallthroughStatement(t *testing.T) {
	result := fallthroughStatement(999)
	if result != "pass fallthrough" {
		t.Errorf("Error")
	} else if result = fallthroughStatement(1); result != "pass fallthrough" {
		t.Errorf("Error on pass fallthrough with i=1")
	} else if result = fallthroughStatement(3); result != "Default value" {
		t.Errorf("Error: expect Default value")
	}
}

func TestSwitchWithNoCondition(t *testing.T) {
	result := switchWithNoCondition(1)
	if result != 2 {
		t.Errorf("Expect %[1]d got %[2]d", 2, result)
	} else if result = switchWithNoCondition(2); result != 2 {
		t.Errorf("Expect %[1]d got %[2]d", 2, result)
	}
}

func TestAllNumbersDiff(t *testing.T) {
	result := allNumbersDiff(237)
	if result != "YES" {
		t.Errorf("Error: expect YES, but got NO")
	} else if result = allNumbersDiff(222); result != "NO" {
		t.Errorf("Error: expect NO, but got YES")
	}
}

func TestHappyTicket(t *testing.T) {
	result := happyTicket(613244)
	if result != "YES" {
		t.Errorf("Expect YES got NO")
	} else if result = happyTicket(123456); result != "NO" {
		t.Errorf("Expect NO got YES")
	}
}

func TestLeapYearCheck(t *testing.T) {
	result := leapYearCheck(2000)
	if result != "YES" {
		t.Errorf("Expect YES got NO")
	} else if result = leapYearCheck(2001); result != "NO" {
		t.Errorf("Expect NO got YES")
	}
}
