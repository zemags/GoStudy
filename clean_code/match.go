package main

import (
	"regexp"
)

var specials = makeSpecials([]rune{
	'\\', '.', '?', '+', '*', '^', '$',
	'|', '{', '}', '[', ']', '(', ')',
})

func match(p, s string) (bool, error) {
	pat := translate(p)
	re, err := regexp.Compile(pat)
	if err != nil {
		return false, err
	}
	isMatch := re.MatchString(s)
	return isMatch, nil
}

func translate(pattern string) string {
	rePat := make([]rune, 0, len(pattern))
	for _, char := range pattern {
		switch char {
		case '*':
			rePat = append(rePat, '.', '*')
		case '?':
			rePat = append(rePat, '.')
		default:
			rePat = append(rePat, escape(char)...)
		}
	}
	return string(rePat)
}

func escape(char rune) []rune {
	if _, isSpecial := specials[char]; isSpecial {
		return []rune{'\\', char}
	}
	return []rune{char}
}

func makeSpecials(chars []rune) map[rune]bool {
	specials := make(map[rune]bool, len(chars))
	for _, char := range chars {
		specials[char] = true
	}
	return specials
}
