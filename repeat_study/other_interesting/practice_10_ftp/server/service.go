package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var BASEDIR = os.Getenv("BASEDIR")

type Cmd struct {
	command string
	arg     string
}

// Receiver
func Receiver(conn net.Conn) {
	defer conn.Close()

	result := "command not found"

	c := cmd(conn)
	if c.command == "cd" {
		result = c.cd()
	} else if c.command == "ls" {
		result = c.ls()
	} else if c.command == "get" {
		result = c.get()
	} else if c.command == "pwd" {
		result = c.pwd()
	} else if c.command == "close" {
		conn.Close()
	}
	conn.Write([]byte("hello " + result))
}

func cmd(conn net.Conn) *Cmd {
	c := &Cmd{}

	reader, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Close()
		return c
	}
	readerSlice := strings.Split(string(reader), " ")
	if len(readerSlice) == 0 {
		// empty command recevied
		conn.Close()
		return c
	}

	c.command = strings.TrimSuffix(readerSlice[0], "\n")
	if len(readerSlice) > 1 {
		c.arg = strings.TrimSuffix(readerSlice[1], "\n")
	}
	return c
}

// pwd return current path
func (c *Cmd) pwd() string {
	currentFolder, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentFolder
}

// ls return content of current folder
func (c *Cmd) ls() (res string) {
	files, err := ioutil.ReadDir(c.pwd())
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fi, err := os.Stat(fmt.Sprintf("%v/%v", c.pwd(), f))
		if err != nil {
			log.Print(err)
			continue
		}
		res += fmt.Sprintf("%v..\t%v\t%v\n", fi.Name()[:10], fi.Size(), fi.ModTime())
	}
	return res
}

// get file content
func (c *Cmd) get() string {
	file, err := os.Open(fmt.Sprintf("%v/%v", c.pwd(), c.arg))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
	}

	return string(content)
}

// cd
func (c *Cmd) cd() string {
	if os.Chdir(c.arg) != nil {
		return "folder not found"
	}
	return "ok"
}
