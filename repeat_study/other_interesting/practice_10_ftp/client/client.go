package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Print(err)
	}
	defer con.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	for {
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(con, text+"\n")
		msg, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Print(">>: " + msg)
	}

}
