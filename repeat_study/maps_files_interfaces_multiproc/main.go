package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	const N = 20

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(">", c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
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

//gorutines and channels
func getFromChannel(ch chan int8, n int8) {
	ch <- n + 1
}

func fiveTimeStringChannel(ch chan string, s string) {
	for i := 0; i < 5; i++ {
		ch <- s + " "
	}
	close(ch)
}

func inputValues(inputStream chan string) {
	for _, v := range []string{"1", "1", "2", "3", "3", "4", "4", "5", "6"} {
		inputStream <- v
	}
	close(inputStream)
}

func removeDuplicates(inputStream, outputStream chan string) {
	lastValue := ""
	for v := range inputStream {
		if v != lastValue {
			outputStream <- v
			lastValue = v
		}
	}
	close(outputStream)
}

func syncChannel() struct{} {
	channel := make(chan struct{}) // struct{} doesnt need memory

	//create anonymous function and make it goroutine
	go func(ch chan struct{}) {
		func() {
			// empty function
			ch <- struct{}{}
		}()
		close(ch)
	}(channel) // pass channel to anonymous function

	return <-channel // read from channel to nowhere
}

func workWithWaitGroup(ch chan uint8) chan uint8 {
	wg := sync.WaitGroup{}
	wg.Add(2) // added number of goroutines

	for i := 0; i < 2; i++ { // call goroutines in for loop
		go func(waitgroup *sync.WaitGroup, channel chan uint8) {
			defer waitgroup.Done() // current goroutine finished
			channel <- 1
		}(&wg, ch)
	}
	wg.Wait() // waiting all goroutines stopped
	close(ch)
	return ch
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type mySyncMap struct {
	mx sync.Mutex
	m  map[int]int
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	wg := &sync.WaitGroup{}
	myMap1 := &mySyncMap{
		m: make(map[int]int),
	}
	myMap2 := &mySyncMap{
		m: make(map[int]int),
	}

	wg.Add(2 * n)
	for i := 0; i < n; i++ {
		go worker(fn, in1, myMap1, i)
		go worker(fn, in2, myMap2, i)
		wg.Add(-2)
	}

	go func() {
		wg.Wait()
		for i := 0; i < n; i++ {
			temp1 := myMap1.Load(i)
			temp2 := myMap2.Load(i)
			result := temp1 + temp2
			out <- result
		}
	}()
}

func worker(fn func(int) int, ch <-chan int, myMap *mySyncMap, key int) {
	res := <-ch
	res = fn(res)
	fmt.Println(res)
	myMap.Save(key, res)
}

func (sm *mySyncMap) Load(key int) int {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	val, _ := sm.m[key]
	return val
}

func (sm *mySyncMap) Save(key, res int) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	sm.m[key] = res
}
