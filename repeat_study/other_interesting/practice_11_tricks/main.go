package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	// commandLineGoDoc()
	methodExpression()
}

func testTables() {
	var peopleTests = []struct {
		name string
		age  int
	}{
		{"Bob", 22},
		{"Ollie", 33},
	}
	fmt.Printf("%s's age %d", peopleTests[0].name, peopleTests[1].age)
}

func embededLock() {
	var hits struct {
		sync.Mutex
		n int
	}

	hits.Lock()
	hits.n++
	hits.Unlock()
}

func commandLineGoDoc() {
	// execute 'go doc -src sync Mutex' command to see Mutex docs
	// get 'go' executable path
	goExecutable, err := exec.LookPath("go")
	if err != nil {
		panic(err)
	}

	// 'go doc -src sync Mutex'
	cmdGoDoc := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "doc", "-src", "sync", "Mutex"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fmt.Println(cmdGoDoc.String())
	// run command
	if err := cmdGoDoc.Run(); err != nil {
		fmt.Println("erorr", err)
	}
}

type People struct {
	name string
	age  int
}

func (p People) HelloValue(s string) {
	fmt.Printf("hello %s from %s\n", p.name, s)
}

func (p *People) HelloPointer(s string) {
	fmt.Printf("hello %s from %s\n", p.name, s)
}

func methodExpression() {
	var p People
	p.HelloValue("John")

	People.HelloValue(p, "Kate")

	(People).HelloValue(p, "Amy")

	f1 := People.HelloValue
	f1(p, "Hugo")

	f2 := (People).HelloValue
	f2(p, "Spam")
}

type Customer struct {
	Name string
	Age  int
}

func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(js)
	return err
}

func savePerson() {
	c := &Customer{Name: "Bob", Age: 22}

	var buf *bytes.Buffer
	if err := c.WriteJSON(buf); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
