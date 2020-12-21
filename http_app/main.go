package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999) //[]byte - type of a slice, 99999 - space for
	//number of elements in slices can will be write
	response.Body.Read(bs) // pass html body to Read interface
	// with our empty byteslice to fill it
	fmt.Println(string(bs))
}
