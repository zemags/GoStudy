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

		} else if !ifSymbolDigit {
			result.WriteString(string(symbol))

		} else if ifSymbolDigit && !unicode.IsDigit(sToRune[idx-1]) && !unicode.IsDigit(sToRune[idx+1]) {
			repeatTime, err := strconv.Atoi(string(symbol))
			if err != nil {
				return "", fmt.Errorf("cannot convert string to int: %v", err)
			}

			if repeatTime == 0 {
				continue
			} else {
				result.WriteString(strings.Repeat(string(sToRune[idx-1]), repeatTime-1))
			}

		} else {
			return "", fmt.Errorf("unsupported input format for string: %s", s)
		}

	}
	return result.String(), nil
}
