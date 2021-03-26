package main

import (
	"fmt"
	"log"
	"net"
)

type Cmd struct {
	command string
}

// Recevier
func (c Cmd) Recevier(conn net.Conn) string {
	switch {
	case c.command == "cd":
		return c.cd()
	case c.command == "ls":
		return c.ls()
	case c.command == "get":
		return c.get()
	case c.command == "close":
		return c.close(conn)
	default:
		return fmt.Sprintf("command %s not found", c.command)
	}
}

// ls return conent of current folder
func (c Cmd) ls() string {
	return ""
}

// get
func (c Cmd) get() string {
	return ""
}

// close
func (c Cmd) close(conn net.Conn) string {
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	return "Close connection"
}

// cd
func (c Cmd) cd() string {
	return ""
}
