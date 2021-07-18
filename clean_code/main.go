package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pattern, src, err := readInput()
	if err != nil {
		fail(err)
	}
	isMatch, err := match(pattern, src)
	if err != nil {
		fail(err)
	}
	if !isMatch {
		os.Exit(1)
	}
	fmt.Println(src)
}

func readInput() (p, s string, err error) {
	flag.StringVar(&p, "p", "", "pattern to match against")
	flag.Parse()
	if p == "" {
		return p, s, errors.New("missing pattern")
	}
	s = strings.Join(flag.Args(), "")
	if s == "" {
		return p, s, errors.New("missing string to match")
	}
	return p, s, nil
}

func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
