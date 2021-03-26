package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	lsn, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// listen port
	conn, err := lsn.Accept()

	for {
		cmd, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		c := Cmd{
			command: cmd,
		}
		c.Recevier(conn)
	}
}
