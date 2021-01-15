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

func TestIsStartsWithUpper(t *testing.T) {
	result := isStartsWithUpper("Быть или не быть.")
	if result != "Right" {
		t.Errorf("expected Right instead got %s", result)
	}
	result = isStartsWithUpper("или не быть.")
	if result != "Wrong" {
		t.Errorf("expected Wrong instead got %s", result)
	}
	result = isStartsWithUpper("Быть или не быть")
	if result != "Wrong" {
		t.Errorf("expected Wrong instead got %s", result)
	}
}

func TestPalindrom(t *testing.T) {
	result := palindrom("топот")
	if result != "Палиндром" {
		t.Errorf("expected Палиндром instead got %s", result)
	}
	result = palindrom("топ")
	if result != "Нет" {
		t.Errorf("expected Нет instead got %s", result)
	}
}

func TestPartOfString(t *testing.T) {
	result := partOfString("hello", "he")
	if result != 0 {
		t.Errorf("expected 0 instead got %d", result)
	}
	result = partOfString("hello", "xy")
	if result != -1 {
		t.Errorf("expected -1 instead got %d", result)
	}
}

func TestOddLetters(t *testing.T) {
	result := oddLetters("ihgewlqlkot")
	if result != "hello" {
		t.Errorf("expected hello instead got %s", result)
	}
}

func TestRemoveMoreThanOneLetter(t *testing.T) {
	result := removeMoreThanOneLetter("zaabcbd")
	if result != "zcd" {
		t.Errorf("expected zcd instead got %s", result)
	}
}

func TestCheckPassword(t *testing.T) {
	result := checkPassword("fdsghdfgjsdDD1")
	if result != "Ok" {
		t.Errorf("expected Ok instead got %s", result)
	}
	result = checkPassword("fdsghdfgjsdDD")
	if result != "Wrong password" {
		t.Errorf("expected Wrong password instead got %s", result)
	}
	result = checkPassword("fds")
	if result != "Wrong password" {
		t.Errorf("expected Wrong password instead got %s", result)
	}
	result = checkPassword("fhjG*J424")
	if result != "Wrong password" {
		t.Errorf("expected Wrong password instead got %s", result)
	}
}

func TestDivide(t *testing.T) {
	result := divide(4, 2)
	if result != 2 {
		t.Errorf("expected 2 instead got %d", result)
	}

	defer func() {
		recover()
	}()

	divide(1, 0)

	t.Error("did not panic")
}

func TestPutStartBetweenLetters(t *testing.T) {
	result := putStartBetweenLetters("abc")
	if result != "a*b*c" {
		t.Errorf("expected a*b*c instead got %s", result)
	}
}

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
