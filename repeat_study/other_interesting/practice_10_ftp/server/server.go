package main

import (
	"log"
	"net"
	"os"
)

func main() {
	lsn, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	for {
		// listen port
		conn, err := lsn.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go Receiver(conn)
	}
}
