package main

import "fmt"

func printSlice() {
	numberFour := 4
	var numberFive int = 5
	numbersSlice := []int{1, 2, 3, numberFour, numberFive}

	for _, number := range numbersSlice {
		fmt.Println("Hello ", number)
		break
	}
}

type birthDate []int

func (bD birthDate) formatPrint() {
	fmt.Println(bD[0]) // will print 12345
}
