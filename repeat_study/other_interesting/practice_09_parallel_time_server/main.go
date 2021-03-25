package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// go run main.go
	// netcat localhost 8000
	lsn, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := lsn.Accept()
		if err != nil {
			log.Print(err) // maybe line break up
			continue
		}
		go handler(con) // for many clients
	}

}

func handler(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client exit
		}
		time.Sleep(1 * time.Second)
	}
}
