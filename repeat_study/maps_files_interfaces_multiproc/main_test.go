package main

import (
	"bytes"
	"testing"
	"time"
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
	numOfBytes, result := sumFromStdin(buf)
	if result != 13 {
		t.Errorf("expected 13 instead got %d", result)
	}
	if numOfBytes != 2 {
		t.Errorf("expected 2 instead got %d", numOfBytes)
	}
}

func TestReadDirsAndFiles(t *testing.T) {
	t.Skip()
	result := readDirsAndFiles("")
	if result != "42" {
		t.Errorf("expected 42 instead got %s", result)
	}
}

func TestFindZero(t *testing.T) {
	t.Skip()
	result := findZero("")
	if result != 10527 {
		t.Errorf("expected 10527 instead got %d", result)
	}
}

func TestCalcMeanPointForGroup(t *testing.T) {
	t.Skip()
	result := calcMeanPointForGroup("")
	if result != 3 {
		t.Errorf("expected 3 instead got %f", result)
	}
}

func TestGetJSONFieldsSum(t *testing.T) {
	t.Skip()
	result := getJSONFieldsSum("")
	if result != 763804981288 {
		t.Errorf("expected 763804981288 instead got %d", result)
	}
}

func TestTimeToString(t *testing.T) {
	t.Skip()
	result := timeToString("1986-04-16T05:20:00+06:00")
	if result != "Wed Apr 16 05:20:00 +0006 1986" {
		t.Errorf("expected Wed Apr 16 05:20:00 +06 1986 instead got %s", result)
	}
}

func TestAddDayToTime(t *testing.T) {
	result := addDayToTime("2020-05-15 08:00:00")
	if result != "2020-05-15 08:00:00" {
		t.Errorf("expected 2020-05-15 08:00:00 instead got %s", result)
	}
	result = addDayToTime("2020-05-15 13:01:00")
	if result != "2020-05-16 13:01:00" {
		t.Errorf("expected 2020-05-16 13:01:00 instead got %s", result)
	}
}

func TestDurBtwnTwoDates(t *testing.T) {
	result := durBtwnTwoDates("13.03.2018 14:00:15,12.03.2018 14:00:15")
	if result != "24h0m0s" {
		t.Errorf("expected 24h0m0s instead got %s", result)
	}

	result = durBtwnTwoDates("12.03.2018 14:00:15,13.03.2018 14:00:15")
	if result != "24h0m0s" {
		t.Errorf("expected 24h0m0s instead got %s", result)
	}
}

func TestParseDate(t *testing.T) {
	result := parseDate("13.03.2018 14:00:15", "02.01.2006 15:04:05")
	if result != time.Date(2018, time.March, 13, 14, 0, 15, 0, time.UTC) {
		t.Errorf("expected 2018-03-13 14:00:15 +0000 UTC instead got %v", result)
	}
}

func TestUnixAddPeriod(t *testing.T) {
	result := unixAddPeriod("12 мин. 13 сек.")
	if result != "Fri May 15 19:28:18 UTC 2020" {
		t.Errorf("expected Fri May 15 19:28:18 UTC 2020 instead got %v", result)
	}
	result = unixAddPeriod("80 мин. 15 сек.")
	if result != "Fri May 15 20:36:20 UTC 2020" {
		t.Errorf("expected Fri May 15 19:28:18 UTC 2020 instead got %v", result)
	}
}

func TestGetFromChannel(t *testing.T) {
	var result int8
	channel := make(chan int8, 1)
	getFromChannel(channel, 1)
	result = <-channel
	close(channel)
	if result != 2 {
		t.Errorf("expect 2 got %d", result)
	}
}

func TestFiveTimeStringChannel(t *testing.T) {
	var result string
	channel := make(chan string, 5)

	go fiveTimeStringChannel(channel, "a")

	for val := range channel {
		result += val
	}

	if result != "a a a a a " {
		t.Errorf("expect aaaaa  got %s", result)
	}
}
func TestInputValues(t *testing.T) {
	var result string
	channel := make(chan string)
	go inputValues(channel)
	for val := range channel {
		result += val
	}
	if result != "112334456" {
		t.Errorf("expect aaaaa  got %s", result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var result string
	inputStream := make(chan string)
	outputStream := make(chan string)
	go inputValues(inputStream)
	go removeDuplicates(inputStream, outputStream)

	for val := range outputStream {
		result += val
	}
	if result != "123456" {
		t.Errorf("expect aaaaa  got %s", result)
	}
}
