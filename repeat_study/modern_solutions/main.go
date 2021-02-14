package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// itPanic()
	// periodicSend()
	reverseString("abc")

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
