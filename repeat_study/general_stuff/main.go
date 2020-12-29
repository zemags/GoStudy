package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int // default value 0 for int
	var x = 10
	var y int = 50

	var (
		variable int    = 99
		sentence string = "Hello"
	)

	var lst1, lst2 []int

	y = y / x
	a = 5 + y
	a++ // a = a + 1

	lst1 = appendToList(lst1, a)
	lst1 = appendToList(lst1, b)
	lst1 = appendToList(lst1, c)

	lst2 = appendToList(lst2, 0)

	fmt.Println(lst1, lst2)

	fmt.Println(variable, sentence)
}

func appendToList(lst []int, value int) []int {
	lst = append(lst, value+1)
	return lst
}

const oneGrad int = 2 // constant value can't be changed
func gradToHours() {
	var d int
	fmt.Scan(&d) // read from input
	h := d / 30
	m := oneGrad * (d % 30)
	fmt.Printf("It is %[1]d hours %[2]d minutes.", h, m)
}

func studyConst() {
	const (
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)
	fmt.Println(c0, c1, c2) // вывод: 0 1 2

	const (
		Sunday = iota // iota its a Go value constant starts with 0 int
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // пропускаем 7
		Add
	)

	fmt.Println(Sunday)   // вывод: 0
	fmt.Println(Saturday) // вывод: 6
	fmt.Println(Add)      // вывод: 8

	// переменные ни в одном блоке const, поэтому индекс не увеличился
	const x = iota // x == 0
	const y = iota // y == 0
}

func shortIfStatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func elseIfStatement(a, b int) bool {
	if a < b {
		return true
	} else if a > b { // like pythons elif
		return false
	}
	return true
}

func switchStatement(i int) string {
	switch i {
	case 0:
		return "Zero"
	case 999:
		return "999"
	default:
		return "Default value"
	}
}

func fallthroughStatement(i int) string {
	switch i {
	case 999:
		fallthrough // the next case will be execute
		//anyway, no matter false or true statement in next case
	case 1:
		return "pass fallthrough"
	default:
		return "Default value"
	}
}

func switchWithNoCondition(i int) int {
	switch { // no condition
	case i == 1:
		return i + 1
	default:
		return i
	}
}

func allNumbersDiff(i int) string {
	a := i / 100     // get first item
	b := i / 10 % 10 // get second item
	c := i % 10      // get third item
	switch {
	case a != b && a != c && b != c:
		return "YES"
	default:
		return "NO"
	}
}

func getFirstDigit(i int) int {
	switch {
	case i == 10000:
		return 1
	case i/1000 > 0:
		return i / 1000
	case i/1000 == 0 && i/100 > 0:
		return i / 100
	case i/1000 == 0 && i/100 == 0 && i/10 > 0:
		return i / 10
	default:
		return i
	}
}

func happyTicket(i int) string {
	first3 := i / 1000
	second3 := i % 1000
	switch {
	case helper(first3) == helper(second3):
		return "YES"
	default:
		return "NO"
	}
}

func helper(i int) int {
	a := i / 100
	b := i / 10 % 10
	c := i % 10
	return a + b + c
}

func leapYearCheck(i int) string {
	switch {
	case i%400 == 0:
		return "YES"
	case i%4 == 0 && i%100 != 0:
		return "YES"
	default:
		return "NO"
	}
}
