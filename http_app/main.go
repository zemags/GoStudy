package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999) //[]byte - type of a slice, 99999 - space for
	// //number of elements in slices can will be write
	// response.Body.Read(bs) // pass html body to Read interface
	// // with our empty byteslice to fill it
	// fmt.Println(string(bs))

	lw := logWriter{}

	// io.Copy(os.Stdout, response.Body)
	io.Copy(lw, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) { // logWrite now implementing writer interface
	fmt.Println(string(bs))
	fmt.Println("BS LEN:", len(bs))
	return len(bs), nil

}
