package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := os.Args[1]

	lsn, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := lsn.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handler(con)
	}
}

func handler(c net.Conn) {
	defer c.Close()
	for {
		t := time.Now()
		zone, offset := t.Zone()
		timeString := fmt.Sprintf("zone:%v, offset: %v, time: %v\n", zone, offset, t.Format("15:04:05"))
		_, err := io.WriteString(c, timeString)
		if err != nil {
			return // client disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
