package main

import "testing"

func TestMinimumFromFour(t *testing.T) {
	result := minimumFromFour(1, 2, 3, 4)
	if result != 1 {
		t.Errorf("expected 1 instead got %d", result)
	}
}

func TestHelperForMin(t *testing.T) {
	result := helperForMin(1, 2, 3, 4)
	if result != true {
		t.Errorf("expected true instead got %t", result)
	}
	result = helperForMin(6, 2, 3, 4)
	if result != false {
		t.Errorf("expected false instead got %t", result)
	}
}

func TestVote(t *testing.T) {
	result := vote([]int{0, 0, 1})
	if result != 0 {
		t.Errorf("expected 0 instead got %d", result)
	}
	result = vote([]int{1, 0, 1})
	if result != 1 {
		t.Errorf("expected 1 instead got %d", result)
	}
}

func TestFibonacci(t *testing.T) {
	result := fibonacci(4)
	if result != 3 {
		t.Errorf("expected 3 instead got %d", result)
	}
}

func TestSumInt(t *testing.T) {
	resLen, resSum := sumInt(1, 2, 3)
	if resLen != 3 && resSum != 6 {
		t.Errorf("expected resLen 3, resSum 6 instead got %d %d", resLen, resSum)
	}
}
