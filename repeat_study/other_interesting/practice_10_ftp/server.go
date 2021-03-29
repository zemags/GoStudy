package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const BASEDIR = "./"

// var ErrInvalidCommand = fmt.Errorf("command not found")

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
		if err != nil {
			log.Print(err)
			continue
		}
		res += fmt.Sprintf("%v..\t%v\t%v\n", fi.Name()[:10], fi.Size(), fi.ModTime())
	}
	return res
}

// get file content
func (c Cmd) get() string {
	commands, err := parse(c.command)
	if err != nil {
		// print error
	}
	filename := commands[1]

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
	if err != nil {
		log.Print(err)
	}

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
	commands, err := parse(c.command)
	if err != nil {
		// print err
	}
	path := commands[1]
	if path == ".." {

	}
	// validate path check with BASEDIR
	// ..
	// if string
	return ""
}

func parse(s string) ([]string, error) {
	if len(s) != 2 {
		return []string{}, errors.New("command not found")
	}
	return strings.Split(s, " "), nil
}
