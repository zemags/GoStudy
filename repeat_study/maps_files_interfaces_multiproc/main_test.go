package main

import (
	"bytes"
	"testing"
)

func TestSimpleCache(t *testing.T) {
	result := simpleCache([]int{1, 2, 3, 2})
	if result != "1 4 9 4" {
		t.Errorf("expected 1 4 9 4 instead got %s", result)
	}
}

func TestWork(t *testing.T) {
	result := work(2)
	if result != 4 {
		t.Errorf("expected 4 instead got %d", result)
	}
}

func TestStringToInt(t *testing.T) {
	result := stringToInt("10abcn")
	if result != 10 {
		t.Errorf("expected 10 instead got %d", result)
	}
}

func TestRemoveOddNumbers(t *testing.T) {
	result := removeOddNumbers(727178)
	if result != 28 {
		t.Errorf("expected 28 instead got %d", result)
	}
	result = removeOddNumbers(301)
	if result != 100 {
		t.Errorf("expected 100 instead got %d", result)
	}
}

func TestMainCalculator(t *testing.T) {
	t.Run("5.0, 5.0, +", testMainCalculator(5.0, 5.0, "+", "10.0000"))
	t.Run("5.0, 5.0, -", testMainCalculator(5.0, 5.0, "-", "0.0000"))
	t.Run("5.0, 5.0, *", testMainCalculator(5.0, 5.0, "*", "25.0000"))
	t.Run("5.0, 5.0, /", testMainCalculator(5.0, 5.0, "/", "1.0000"))
	t.Run("5.0, 5.0, &", testMainCalculator(5.0, 5.0, "&", "неизвестная операция"))
	t.Run("5, 5.0, +", testMainCalculator(5, 5.0, "+", "value=5: int"))
	t.Run("5.0, 5, +", testMainCalculator(5.0, 5, "+", "value=5: int"))

}

func testMainCalculator(x, y interface{}, o interface{}, expected string) func(*testing.T) {
	return func(t *testing.T) {
		result := mainCalculator(x, y, o)
		if result != expected {
			t.Errorf("expected %s instead got %s", expected, result)
		}
	}
}

func TestCalculator(t *testing.T) {
	result := calculator(5.0, 5.0, "+")
	if result != 10.0 {
		t.Errorf("expected 10.0 instead got %.f", result)
	}
}

func Test_getOperationType(t *testing.T) {
	result := _getOperationType("+")
	if result != true {
		t.Errorf("expected true instead got %t", result)
	}
	result = _getOperationType(1)
	if result != false {
		t.Errorf("expected false instead got %t", result)
	}
}

func Test__getType(t *testing.T) {
	result := _getType(5)
	if result != false {
		t.Errorf("expected false instead got %t", result)
	}
	result = _getType(5.0)
	if result != true {
		t.Errorf("expected true instead got %t", result)
	}
}

func TestFillBattery(t *testing.T) {
	result := fillBattery("1000010011")
	if result != "[      XXXX]" {
		t.Errorf("expected [      XXXX] instead got %s", result)
	}
}

func TestSumFromStdin(t *testing.T) {
	buf := bytes.NewBufferString("1\n2\n10")
	result := sumFromStdin(buf)
	if result != 2 {
		t.Errorf("expected 2 instead got %d", result)
	}
}
