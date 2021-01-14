package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println(minimumFromFour(4, 5, 6, 7))
	// multiplyPointersVals(2, 5)
	fmt.Println(removeMoreThanOneLetter("zaabcbd"))
}

func minimumFromFour(a, b, c, d int) int {
	if helperForMin(a, b, c, d) {
		return a
	} else if helperForMin(b, a, c, d) {
		return b
	} else if helperForMin(c, a, b, d) {
		return c
	} else {
		return d
	}
}

func helperForMin(g, x, y, z int) bool {
	if g <= x && g <= y && g <= z {
		return true
	}
	return false
}

func vote(a []int) int {
	var zeros, ones int
	for _, i := range a {
		if i == 0 {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return 0
	}
	return 1
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func sumInt(a ...int) (int, sum int) { // ..переменное число аргументов
	for _, elem := range a {
		sum += elem
	}
	return len(a), sum
}

func multiplyPointersVals(x1 *int, x2 *int) int {
	x := *x1 * *x2
	return x
}

func shiftPointersVals(x1 *int, x2 *int) (int, int) {
	*x1, *x2 = *x2, *x1
	return *x1, *x2
}

// Structs
type Circle struct {
	x float64
	y float64
	r float64
}

func prepareCircle() float64 {
	c := Circle{x: 2, y: 3, r: 10}
	circleArea := c.getCircleArea()
	return circleArea
}

func (c *Circle) getCircleArea() float64 {
	return 3.14 * c.r * c.r
}

type Android struct {
	On    bool
	Ammo  int
	Power int
}

func (a *Android) Shoot() bool {
	if a.On {
		if a.Ammo > 0 {
			(*a).Ammo--
			return true
		}
	}
	return false
}

func (a *Android) RideBike() bool {
	if a.On {
		if a.Power > 0 {
			(*a).Power--
			return true
		}
	}
	return false
}

// strings
func isStartsWithUpper(text string) string {
	text = strings.Trim(text, "\n\r")
	runeText := []rune(text)

	if unicode.IsUpper(runeText[0]) && string(runeText[len(runeText)-1]) == "." {
		return "Right"
	}
	return "Wrong"
}

func palindrom(text string) string {
	var reverseRuneText []string

	runeText := []rune(strings.Trim(text, "\n\r"))

	for i := len(runeText) - 1; i != -1; i-- {
		reverseRuneText = append(reverseRuneText, string(runeText[i]))
	}
	if text == strings.Join(reverseRuneText, "") {
		return "Палиндром"
	}
	return "Нет"
}

func partOfString(x, s string) int {
	s = strings.Trim(s, "\n\r")
	x = strings.Trim(x, "\n\r")
	if strings.Count(x, s) > 0 {
		return strings.Index(x, s)
	}
	return -1
}

func oddLetters(text string) string {
	var result string
	for idx, letter := range text { // in loop letter its rune!
		if idx%2 != 0 {
			result += string(letter)
		}
	}
	return result
}

func removeMoreThanOneLetter(text string) string {
	copyText := []rune(text)
	for _, letter := range copyText {
		if strings.Count(text, string(letter)) >= 2 {
			text = strings.Replace(text, string(letter), "", -1)
		}
	}
	return text
}

func checkPassword(text string) string {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	const digit = "1234567890"
	flag := true
	var isDigit, isLetter = false, false
	if len(text) >= 5 {
		for _, r := range text {
			if !strings.Contains(alpha, strings.ToLower(string(r))) && !strings.Contains(digit, strings.ToLower(string(r))) {
				flag = false
			}
			if strings.Contains(alpha, strings.ToLower(string(r))) {
				isLetter = true
			}
			if strings.Contains(digit, strings.ToLower(string(r))) {
				isDigit = true
			}
		}
		if flag && isLetter && isDigit {
			return "Ok"
		}
	}
	return "Wrong password"
}

//errors panic
func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func putStartBetweenLetters(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			res += string(s[i])
		} else {
			res += string(s[i]) + "*"
		}
	}
	return res
}

//maps
func simpleCache(s []int) string {
	m := make(map[int]int)
	result := ""
	for _, a := range s {
		if value, ok := m[a]; ok {
			m[a] = value
		} else {
			m[a] = work(a)
		}
		result += fmt.Sprintf("%d ", m[a])
	}
	return strings.Trim(result, " ")
}

func work(i int) int {
	return i * i
}
