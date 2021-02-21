package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// practice context

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunciton := context.WithCancel(ctx)

	// defer func clean all resources for this and child context
	defer func() {
		fmt.Println("canceling context from main defer")
		cancelFunciton()
	}()
	// context canceling after calling random timeout and all child
	// contexts need to stop
	go func() {
		sleepRandom("Main", nil)
		cancelFunciton()
		fmt.Println("main sleep complete, canceling context")
	}()
	// working
	doWorkContext(ctxWithCancel)
}

// slow work func with context
func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	sleeptimeChan := make(chan int)

	// run slow func in goroutine
	go sleepRandom("sleepRandomContext", sleeptimeChan)

	// exit when context life is over
	select {
	case <-ctx.Done():
		// clean ram, signal(by channel) to all goroutines to stop
		// its happend if timeout is over or doWorkContext or main
		// call cancelFunction
		fmt.Println("time to return all work is over")
	case sleeptime := <-sleeptimeChan:
		// the work stop before context die or cancel
		fmt.Println("slept for ", sleeptime)
	}
}

// slow func
func sleepRandom(fromFunction string, ch chan int) {
	defer func() {
		fmt.Println(fromFunction, "sleepRandom complete")
	}()
	// imitate slow work - sleep for random time
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100

	fmt.Println(fromFunction, "start sleep", sleeptime)
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "complete sleep", sleeptime)

	// write to channel if it passed
	if ch != nil {
		ch <- sleeptime
	}
}

// helper func
// ctx in case its a context.WithCancel
func doWorkContext(ctx context.Context) {
	// all child context are over after 150 ms
	timeout := time.Duration(150) * time.Millisecond
	// ctxWithTimeout its a contex.WithCancel child
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, timeout)

	// cancel func for clean resources after doWorkContext finished
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	// select for exit when context is over
	select {
	case <-ctx.Done():
		// when context finished (when call cancelFunction from main)
		fmt.Println("doWrokContext return")
	case <-ch:
		// here if work is finished before context done
		fmt.Println("sleepRandomContext returned")
	}
}
