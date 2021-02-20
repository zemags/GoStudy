package main

import (
	"context"
	"fmt"
)

// practise context

func main() {

}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	// slow work func
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	sleeptimeChan := make(chan int)

	// run slow func in goroutine
	go sleepRandom("sleepRandomConext", sleeptimeChan)

	// exit when context life is over
	select {
	case <-ctx.Done():
		// clean ram, signal(by channel) to all goroutines to stop
		fmt.Println("time to return all work is over")
	case sleeptime := <-sleeptimeChan:
		// the wowk stop before context die or cancel
		fmt.Println("slept for ", sleeptime, "ms")
	}
}
