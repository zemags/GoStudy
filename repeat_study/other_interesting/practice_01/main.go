package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// unpackString()
}

func unpackString(s string) (string, error) {
	sToRune := []rune(s)
	var result strings.Builder
	prevIsDigit := false
	repeatTime := 0
	for idx, symbol := range sToRune {
		ifSymbolDigit := unicode.IsDigit(symbol)
		if ifSymbolDigit && idx == 0 {
			return "", fmt.Errorf("unsupported input format for string: %s", s)
		}
		if idx != len(s) {

			if !ifSymbolDigit {

				if prevIsDigit {
					return "", fmt.Errorf("unsupported input format for string: %s", s)

				} else if unicode.IsDigit(sToRune[idx+1]) && !prevIsDigit {
					repeatTime, _ = strconv.Atoi(string(sToRune[idx+1]))
				} else {
					repeatTime = 1
				}
				result.WriteString(strings.Repeat(string(symbol), repeatTime))
				fmt.Println(string(symbol), repeatTime, result.String())
				prevIsDigit = true
			}
		} else {
			result.WriteString((string(symbol)))
			prevIsDigit = false
		}
	}
	return result.String(), nil
}
