package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args[0] // temp path to compiled go exe
	fileName := os.Args[1]         // first command line argument
	file, err := os.Open(fileName) // return a pointer to a file
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// } else {
	// 	scanner := bufio.NewScanner(file)
	// 	scanner.Split(bufio.ScanLines) // ScanLines as an input to the method Split
	// 	var text []string
	// 	for scanner.Scan() {
	// 		text = append(text, scanner.Text())
	// 	}
	// 	file.Close()
	// 	fmt.Println(text) // will print [There is a not empty file.]
	// }

	// original solution
	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
