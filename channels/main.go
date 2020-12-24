package main

import (
	"fmt"
	"net/http"
)

func main() { // main routine, without channels will not be wating other child
	// routines to finished
	links := []string{
		"http://google.com",
		"http://yandex.ru",
		"http://mail.ru",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // create channel 'c' that recivied type string

	for _, link := range links {
		go checkLink(link, c) // 'go' - its go routine, pass 'c' - channel
		// channels need to communicate between go routines, and channel accept
		// only one type like string or int or smth
	}

	// fmt.Println(<-c) // pass channel recivied value into print
	// fmt.Println(<-c) // also wait some data
	// or do it like below
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // value coming from 'c' is a blocking call
	}
}

func checkLink(link string, c chan string) { // 'c' -variable of type 'chan'-channel
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "DOWN")
		c <- "Might be down i think" // pass value of type string to channel
		return
	}

	fmt.Println(link, "UP")
	c <- "Its up"
}
