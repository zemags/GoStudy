package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour(4, 5, 6, 7))
	// multiplyPointersVals(2, 5)
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
