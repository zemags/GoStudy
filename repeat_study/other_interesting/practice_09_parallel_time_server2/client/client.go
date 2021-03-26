package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// define client
	port := os.Args[1]
	con, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()
	copyFromSrc(os.Stdout, con)
}

func copyFromSrc(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
