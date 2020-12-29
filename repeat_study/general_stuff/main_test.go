package main

import (
	"fmt"
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
		t.Errorf("Values doesnt match, expected %[1]f, but get %[2]f", expectValue, result)
	}
}

func TestElseIfStatement(t *testing.T) {
	result := elseIfStatement(2, 4)
	if !result {
		t.Errorf("expected true")
	}

	result = elseIfStatement(4, 3)
	if result {
		t.Errorf("expected false")
	}

	result = elseIfStatement(1, 1)
	if !result {
		t.Errorf("expected true")
	}
}

func TestSwitchStatement(t *testing.T) {
	result := switchStatement(0)
	if result != "Zero" {
		t.Errorf("")
	}
	result = switchStatement(999)
	if result != "999" {
		t.Errorf("expected 999, got %[1]s", result)
	}
	result = switchStatement(3)
	if result != "Default value" {
		t.Errorf("expected 999, got %[1]s", result)
	}
}

func TestFallthroughStatement(t *testing.T) {
	result := fallthroughStatement(999)
	if result != "pass fallthrough" {
		t.Errorf("Error")
	} else if result = fallthroughStatement(1); result != "pass fallthrough" {
		t.Errorf("Error on pass fallthrough with i=1")
	} else if result = fallthroughStatement(3); result != "Default value" {
		t.Errorf("Error: expected Default value")
	}
}

func TestSwitchWithNoCondition(t *testing.T) {
	result := switchWithNoCondition(1)
	if result != 2 {
		t.Errorf("expected %[1]d got %[2]d", 2, result)
	} else if result = switchWithNoCondition(2); result != 2 {
		t.Errorf("expected %[1]d got %[2]d", 2, result)
	}
}

func TestAllNumbersDiff(t *testing.T) {
	result := allNumbersDiff(237)
	if result != "YES" {
		t.Errorf("Error: expected YES, but got NO")
	} else if result = allNumbersDiff(222); result != "NO" {
		t.Errorf("Error: expected NO, but got YES")
	}
}

func TestHappyTicket(t *testing.T) {
	result := happyTicket(613244)
	if result != "YES" {
		t.Errorf("expected YES got NO")
	} else if result = happyTicket(123456); result != "NO" {
		t.Errorf("expected NO got YES")
	}
}

func TestLeapYearCheck(t *testing.T) {
	result := leapYearCheck(2000)
	if result != "YES" {
		t.Errorf("expected YES got NO")
	} else if result = leapYearCheck(2001); result != "NO" {
		t.Errorf("expected NO got YES")
	}
}

// Multiple Test Cases per Function
func TestGetFirstDigit(t *testing.T) {
	t.Run("10000", testFirstDigit(10000, 1))
	t.Run("2345", testFirstDigit(2345, 2))
	t.Run("345", testFirstDigit(345, 3))
	t.Run("45", testFirstDigit(45, 4))
	t.Run("5", testFirstDigit(5, 5))
}

func testFirstDigit(number int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		result := getFirstDigit(number)
		if result != expected {
			t.Errorf(fmt.Sprintf("Expected the %d but instead got %d", expected, result))
		}
	}
}

func TestSumBetweenTwoNumber(t *testing.T) {
	result := sumBetweenTwoNumber(1, 5)
	if result != 15 {
		t.Errorf("Expected 15 instead got %d", result)
	}
}

func TestSumOfTwoInRange(t *testing.T) {
	result := sumOfTwoInRange([]int{38, 24, 800, 8, 16})
	if result != 40 {
		t.Errorf("Expected 40 instead got %d", result)
	}
}

func TestTwoNumbersDigitCompare(t *testing.T) {
	result := twoNumbersDigitCompare(564, 8954)
	if result != "5 4" {
		t.Errorf("expected 5 4 instead got %s", result)
	}
}

func TestCountElements(t *testing.T) {
	t.Run("1,3,3,1,0", testCountElements([]int{1, 3, 3, 1, 0}, 2))
	t.Run(", 3, 3, 6, 6, 7, 6, 0", testCountElements([]int{1, 3, 3, 6, 6, 7, 6, 0}, 3))
	t.Run("1, 5, 5, 5, 0", testCountElements([]int{1, 5, 5, 5, 0}, 3))
	t.Run("5, 1, 5, 5, 0}", testCountElements([]int{5, 1, 5, 5, 0}, 3))
}

func testCountElements(s []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		result := countElements(s)
		if result != expected {
			t.Errorf("expected %d instead got %d", expected, result)
		}
	}
}
