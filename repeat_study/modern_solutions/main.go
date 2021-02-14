package main

import (
	"fmt"
	"time"
)

func main() {
	itPanic()
	// periodicSend()

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

func reverseString(s string) string {
	fmt.Printf("Reverse string: %s", s)
	return ""
}

func periodicSend() {

	child := func(c chan int) {
		i := 0
		for i <= 3 {
			c <- i
			i++
			time.Sleep(1 * time.Second)
		}
		close(c)
	}

	ch := make(chan int)
	go child(ch)
	for i := range ch {
		fmt.Println(i)
	}
	_, ok := <-ch
	fmt.Println(ok)
}
