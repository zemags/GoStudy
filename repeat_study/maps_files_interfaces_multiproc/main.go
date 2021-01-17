package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	work(2)
}

//maps
func simpleCache(s []int) string {
	m := make(map[int]int)
	result := ""
	for _, a := range s {
		if value, ok := m[a]; ok {
			m[a] = value
		} else {
			m[a] = work(a)
		}
		result += fmt.Sprintf("%d ", m[a])
	}
	return strings.Trim(result, " ")
}

func work(i int) int {
	return i * i
}

// types convertion
func stringToInt(a string) int64 {
	var stringResult string
	var intResult int
	for _, letter := range strings.Split(a, "") {
		_, err := strconv.Atoi(letter)
		if err == nil {
			stringResult += string(letter)
		}
	}
	intResult, err := strconv.Atoi(stringResult)
	if err != nil {
		panic(err)
	}
	return int64(intResult)
}

//lambda functions
func removeOddNumbers(i uint) uint {
	var res int
	fn := func(x uint) string {
		var initial uint = 10000000
		var result string
		for i := 0; i < 8; i++ {
			current := x / initial
			if current > 0 {
				x = x - current*initial
			}
			initial = initial / 10
			if current != 0 {
				if current%2 == 0 {
					result += strconv.FormatUint(uint64(current), 10)
				}
			}
		}
		if result == "0" || result == "" {
			return "100"
		}
		return result
	}
	res, err := strconv.Atoi(fn(i))
	if err != nil {
		panic(err)
	}
	return uint(res)
}

//interfaces
func _getType(val interface{}) bool {
	if _, isFloat64 := val.(float64); isFloat64 {
		return true
	}
	return false
}

func _getOperationType(val interface{}) bool {
	if _, isString := val.(string); isString && (val.(string) == "+" || val.(string) == "-" || val.(string) == "*" || val.(string) == "/") {
		return true
	}
	return false
}

func calculator(value1, value2 interface{}, operator interface{}) float64 {
	var result float64
	if operator.(string) == "+" {
		result = value1.(float64) + value2.(float64)
	} else if operator.(string) == "-" {
		result = value1.(float64) - value2.(float64)
	} else if operator.(string) == "*" {
		result = value1.(float64) * value2.(float64)
	} else {
		if value2.(float64) != 0 {
			result = value1.(float64) / value2.(float64)
		} else {
			panic("Division by zero")
		}
	}
	return result
}

func mainCalculator(value1, value2 interface{}, operation interface{}) string {
	var stringResult string
	if !_getType(value1) {
		stringResult = fmt.Sprintf("value=%v: %T", value1, value1)
	} else if !_getType(value2) {
		stringResult = fmt.Sprintf("value=%v: %T", value2, value2)
	} else if !_getOperationType(operation) {
		stringResult = fmt.Sprint("неизвестная операция")
	} else if _getType(value1) && _getType(value2) && _getOperationType(operation) {
		result := calculator(value1, value2, operation)
		stringResult = fmt.Sprintf("%.4f", result)
	}
	return stringResult
}

type Battery struct {
	Charge string
}

func (b Battery) String() string {
	count := strings.Count(b.Charge, "1")
	return fmt.Sprintf(
		"[%s%s]",
		strings.Repeat(" ", 10-count),
		strings.Repeat("X", count),
	)
}

func fillBattery(s string) string {
	batteryForTest := Battery{
		Charge: s,
	}
	return fmt.Sprint(batteryForTest)
}

//work with files
func sumFromStdin(stdin io.Reader) (int, int) {
	reader := bufio.NewReader(stdin)
	var result int

	for {
		line, _, err := reader.ReadLine() // without \n = End of line bytes
		if err == io.EOF {
			break
		}
		num, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		result += num
	}

	writer := bufio.NewWriter(os.Stdout)
	numberOfBytes, err := writer.WriteString(strconv.Itoa(result))
	if err != nil {
		panic(err)
	}

	return numberOfBytes, result
}
