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

func TestMultiplyPointersVals(t *testing.T) {
	var x, y int = 2, 2
	var x1 = &x // pointer to value
	var x2 = &y
	result := multiplyPointersVals(x1, x2)
	if result != 4 {
		t.Errorf("expected 4 instead got %d", result)
	}
}

func TestShiftPointersVals(t *testing.T) {
	var x, y int = 1, 2
	var x1 = &x // pointer to value
	var x2 = &y
	res1, res2 := shiftPointersVals(x1, x2)
	if res1 != 2 && res2 != 1 {
		t.Errorf("expected 2, 1 instead got %d, %d", res1, res2)
	}
}

func TestPrepareCircle(t *testing.T) {
	result := prepareCircle()
	if result != 314 {
		t.Errorf("expected 314 instead got %f", result)
	}
}

func TestGetCircleArea(t *testing.T) {
	testCircle := Circle{1, 1, 1}
	result := testCircle.getCircleArea()
	if result != 3.14 {
		t.Errorf("expected 3.14 instead got %f", result)
	}
}

func TestShoot(t *testing.T) {
	testAndroid := Android{true, 1, 1}
	result := testAndroid.Shoot()
	if result != true {
		t.Errorf("expected true instead got %t", result)
	}
	testAndroid = Android{false, 1, 1}
	result = testAndroid.Shoot()
	if result != false {
		t.Errorf("expected false instead got %t", result)
	}
}

func TestRideBike(t *testing.T) {
	testAndroid := Android{true, 1, 1}
	result := testAndroid.RideBike()
	if result != true {
		t.Errorf("expected true instead got %t", result)
	}
	testAndroid = Android{false, 1, 1}
	result = testAndroid.RideBike()
	if result != false {
		t.Errorf("expected false instead got %t", result)
	}
}
