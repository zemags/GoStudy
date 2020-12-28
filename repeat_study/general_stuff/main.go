package main

import "fmt"

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
		Sunday = iota // iota its a Go value starts with 0 int
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
