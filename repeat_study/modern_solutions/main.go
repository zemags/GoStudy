package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// itPanic()
	// periodicSend()
	// reverseString("abc")
	// buffChan()
	selectStatement()
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
	var reverseRuneText []string
	runeText := []rune(s)

	for i := len(runeText) - 1; i != -1; i-- {
		reverseRuneText = append(reverseRuneText, string(runeText[i]))
	}
	fmt.Printf("Reverse string: %s", strings.Join(reverseRuneText, ""))
	return strings.Join(reverseRuneText, "")
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

func buffChan() {
	buffCh := make(chan int, 5)
	buffCh <- 3 //FIFO
	buffCh <- 2
	fmt.Println(<-buffCh) // 3
	//buu channels block when the are either full or empty
}

func waitAndSend(v, i int) chan int {
	outCh := make(chan int)
	child := func(v, i int) {
		time.Sleep(time.Duration(i) * time.Second)
		outCh <- v + 1
	}
	go child(1, 3)
	return outCh
}

func selectStatement() {
	ic := make(chan int)
	select {
	case v1 := <-waitAndSend(3, 2):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 1):
		fmt.Println(v2)
	case ic <- 23:
		fmt.Println(ic)
	default:
		fmt.Println("all chan are slow")
	}
}
