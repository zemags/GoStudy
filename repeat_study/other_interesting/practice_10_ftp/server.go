package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const BASEDIR = "./"

type Cmd struct {
	command string
}

// Recevier
func (c Cmd) Recevier(conn net.Conn) string {
	commandSlice := strings.Split(c.command, " ")
	if len(commandSlice) == 0 {
		return fmt.Sprintf("command %s not found", c.command)
	}
	command := commandSlice[1]
	switch {
	case command == "cd":
		return c.cd()
	case command == "ls":
		return c.ls()
	case command == "get":
		return c.get()
	case command == "close":
		return c.close(conn)
	default:
		return fmt.Sprintf("command %s not found", command)
	}
}

// pwd return current path
func (c Cmd) pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// ls return content of current folder
func (c Cmd) ls() (res string) {
	files, err := ioutil.ReadDir(BASEDIR)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fi, err := os.Stat(fmt.Sprintf("%v/%v", c.pwd(), f))
		res += fmt.Sprintf("%v%v\n", f.Name())
	}
	return res
}

// get file content
func (c Cmd) get() string {
	if len(c.command) != 2 {
		return ""
	}
	filename := strings.Split(c.command, " ")[1]

	file, err := os.Open(fmt.Sprintf("%v/%v", c.pwd(), filename))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	content, err := ioutil.ReadAll(file)

	return string(content)
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
	// ..
	// if string
	return ""
}
