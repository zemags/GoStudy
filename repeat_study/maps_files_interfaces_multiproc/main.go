package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(timeToString("1986-04-16T05:20:00+06:00"))
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

func readDirsAndFiles(root string) string {
	var result string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil && info.IsDir() { // ignore dirs
			return nil // ignore current step
		}

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.Comma = ','

		count := 0
		for {
			record, err := reader.Read()
			if err != nil {
				break
			}
			if len(record) == 10 {
				if count == 4 {
					result = record[2]
				}
				count++
			}
		}
		return nil
	}

	err := filepath.Walk(root, walkFunc)
	if err != nil {
		panic(err)
	}
	return result
}

func findZero(p string) int {
	sliceOfString := []string{}
	var result int

	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 500000)
	line, _, err := reader.ReadLine() // without \n = End of line bytes
	if err != nil {
		panic(err)
	}

	sliceOfString = strings.Split(string(line), ";")

	for idx, item := range sliceOfString {
		if item == "0" {
			result = idx + 1
		}
	}
	return result
}

func calcMeanPointForGroup(p string) float32 {
	type studentInfo struct {
		Rating []int
	}
	type group struct {
		Students []studentInfo
	}

	var (
		jsonStruct  group
		pointsCount int
		result      float32
	)

	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &jsonStruct); err != nil {
		panic(err)
	}
	studentsCount := len(jsonStruct.Students)

	for _, student := range jsonStruct.Students {
		pointsCount += len(student.Rating)
	}

	result = float32(pointsCount) / float32(studentsCount)
	return result
}

func getJSONFieldsSum(p string) int {
	var result int
	type jsonStruct struct {
		GlobalID int `json:"global_id"`
	}

	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	dec.Token() // read open bracket

	for dec.More() {
		var s jsonStruct
		err := dec.Decode(&s)
		if err != nil {
			panic(err)
		}
		result += s.GlobalID
	}
	dec.Token() // read closing bracket
	return result
}

//work with time
func timeToString(s string) string {
	givenTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return givenTime.Format(time.UnixDate)
}

func addDayToTime(s string) string {
	const layout = "2006-01-02 15:04:05" // REFERNCE TIME
	givenTime, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	if givenTime.Hour() < 13 {
		return s
	}
	givenTime = givenTime.AddDate(0, 0, 1)
	return givenTime.Format(layout)
}

func durBtwnTwoDates(s string) string {
	const layout = "02.01.2006 15:04:05"
	sliceString := strings.Split(s, ",")
	firstDate := parseDate(sliceString[0], layout)
	secondDate := parseDate(sliceString[1], layout)

	if firstDate.Before(secondDate) {
		return fmt.Sprintf("%v", secondDate.Sub(firstDate))
	}
	return fmt.Sprintf("%v", firstDate.Sub(secondDate))
}

func parseDate(s string, layout string) time.Time {
	s = strings.Trim(s, " ")
	date, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s, layout, date)
	return date
}

func unixAddPeriod(s string) string {
	const now = 1589570165
	s = strings.Replace(strings.Replace(s, " мин. ", "m", 1), " сек.", "s", 1)
	unixTime, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	time2 := time.Unix(now, 0).UTC()
	return time2.Add(unixTime).Format(time.UnixDate)
}
