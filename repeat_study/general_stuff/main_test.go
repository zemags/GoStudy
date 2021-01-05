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
	t.Run("2, 4", testElseIfStatement(2, 4, true))
	t.Run("2, 4", testElseIfStatement(4, 3, false))
	t.Run("2, 4", testElseIfStatement(1, 1, true))
}

func testElseIfStatement(x int, y int, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		result := elseIfStatement(x, y)
		if result != expected {
			t.Errorf("exptected %v, but got %v", expected, result)
		}
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

func TestCutFloat(t *testing.T) {
	t.Run("-000012.2123", testCutFloat(-000012.2123, "-12.21"))
	t.Run("1000001", testCutFloat(1000001, "1.000001e+06"))
	t.Run("12.12345678", testCutFloat(12.12345678, "146.9782"))
}

func testCutFloat(x float64, expected string) func(*testing.T) {
	return func(t *testing.T) {
		result := cutFloat(x)
		if result != expected {
			t.Errorf("expected %v instead got %v", expected, result)
		}
	}
}

func TestWorkWithArray(t *testing.T) {
	var expectedArray [10]uint8 = [10]uint8{99, 137, 151, 187, 117, 71, 20, 187, 93, 67}
	var workArray [10]uint8 = [10]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67}
	var idxArray [6]uint8 = [6]uint8{1, 2, 3, 5, 7, 8}
	result := workWithArray(workArray, idxArray)

	if result != expectedArray {
		t.Errorf("expected %v instead got %v", expectedArray, result)
	}
}

func TestGet4Element(t *testing.T) {
	var sBig [5]int = [5]int{1, -231, 24, 49, 6}
	result := get4Element(sBig)
	if result != 49 {
		t.Errorf("expected %v instead got %v", 49, result)
	}
}

func TestSumOfTree(t *testing.T) {
	result := sumOfTree(745)
	if result != 16 {
		t.Errorf("expected %v instead got %v", 16, result)
	}
}

func TestReverseNumber(t *testing.T) {
	result := reverseNumber(123)
	if result != "321" {
		t.Errorf("expected %v instead got %v", 321, result)
	}
}

func TestCountHoursAndMinutes(t *testing.T) {
	result := countHoursAndMinutes(13257)
	if result != "3 40" {
		t.Errorf("expected %v instead got %v", 13257, result)
	}
}

func TestIfTriangleRight(t *testing.T) {
	result := ifTriangleRight(6, 8, 10)
	if result != "Прямоугольный" {
		t.Errorf("expected Прямоугольный instead got %s", result)
	}
	result = ifTriangleRight(3, 2, 10)
	if result != "Непрямоугольный" {
		t.Errorf("expected Непрямоугольный instead got %s", result)
	}
}

func TestIfTriangleExist(t *testing.T) {
	result := ifTriangleExist(4, 5, 6)
	if result != "Существует" {
		t.Errorf("expected Существует instead got %s", result)
	}
	result = ifTriangleExist(2, 1, 1)
	if result != "Не существует" {
		t.Errorf("expected Не существует instead got %s", result)
	}
}
