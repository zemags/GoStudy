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
	var (
		result     strings.Builder
		repeatTime int
	)

	for idx, symbol := range sToRune {

		if idx != len(s)-1 {
			ifSymbolDigit := unicode.IsDigit(symbol)
			ifNextSymbolDigit := unicode.IsDigit(sToRune[idx+1])

			if (ifSymbolDigit && idx == 0) || (ifSymbolDigit && ifNextSymbolDigit) {
				return "", fmt.Errorf("unsupported input format for string: %s", s)
			} else if !ifSymbolDigit {
				prevIsDigit := false

				if ifNextSymbolDigit && !prevIsDigit {
					repeatTime, _ = strconv.Atoi(string(sToRune[idx+1]))
					prevIsDigit = true
				} else {
					repeatTime = 1
				}
				result.WriteString(strings.Repeat(string(symbol), repeatTime))
				fmt.Println(string(symbol), repeatTime, result.String())
			}
		} else {
			result.WriteString((string(symbol)))
		}
	}
	return result.String(), nil
}
