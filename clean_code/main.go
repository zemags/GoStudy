package main

import (
	"fmt"

	"github.com/sahilm/fuzzy"
)

func main() {
	pttn, src, err := readInput()
	if err != nil {
		fail(err)
	}
	matches := fuzzy.Find(pttn, []string{src})
	isMatch := len(matches) > 0
	if err != nil {
		fail(err)
	}
	fmt.Println(src)
}

func readInput(p, s string, err error) {

}

func fail(err error) {

}
