package client

import (
	"log"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
}
