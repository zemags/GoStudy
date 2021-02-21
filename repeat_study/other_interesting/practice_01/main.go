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
	for idx, symbol := range sToRune {
		ifSymbolDigit := unicode.IsDigit(symbol)
		if ifSymbolDigit && idx == 0 {
			return "", fmt.Errorf("unsupported input format for string: %s", s)
		} else {
			if unicode.IsDigit(symbol) && !unicode.IsDigit(sToRune[idx-1]) {
				repeatTime, err := strconv.Atoi(string(symbol))
				if err != nil {
					return "", fmt.Errorf("cannot convert string to int: %v", err)
				}
				result.WriteString(strings.Repeat(string(sToRune[idx-1]), repeatTime-1))
			} else {
				result.WriteString(string(symbol))
			}
		}
	}
	return result.String(), nil
}
