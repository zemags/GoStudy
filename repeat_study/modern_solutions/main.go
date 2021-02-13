package main

import "fmt"

func main() {
	itPanic()
}

func itPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //A panic happend
			fmt.Println("Recovered from panic")
		}
	}()
	panic("A panic happend")
}
