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
