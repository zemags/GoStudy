package main

import (
	"fmt"
	"math"
)

func main() {
	// var a, b, c int // default value 0 for int
	// var x = 10
	// var y int = 50

	// var (
	// 	variable int    = 99
	// 	sentence string = "Hello"
	// )

	// var lst1, lst2 []int

	// y = y / x
	// a = 5 + y
	// a++ // a = a + 1

	// lst1 = appendToList(lst1, a)
	// lst1 = appendToList(lst1, b)
	// lst1 = appendToList(lst1, c)

	// lst2 = appendToList(lst2, 0)

	// fmt.Println(lst1, lst2)

	// fmt.Println(variable, sentence)
	// fmt.Println(countElements([]int{1, 3, 3, 1, 0}))
	fmt.Println(placeInFibonacciLine(4))
}

func appendToList(lst []int, value int) []int {
	lst = append(lst, value+1)
	return lst
}

const oneGrad int = 2 // constant value can't be changed
func gradToHours() {
	var d int
	fmt.Scan(&d) // read from input
	h := d / 30
	m := oneGrad * (d % 30)
	fmt.Printf("It is %[1]d hours %[2]d minutes.", h, m)
}

func studyConst() {
	const (
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)
	fmt.Println(c0, c1, c2) // вывод: 0 1 2

	const (
		Sunday = iota // iota its a Go value constant starts with 0 int
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // пропускаем 7
		Add
	)

	fmt.Println(Sunday)   // вывод: 0
	fmt.Println(Saturday) // вывод: 6
	fmt.Println(Add)      // вывод: 8

	// переменные ни в одном блоке const, поэтому индекс не увеличился
	const x = iota // x == 0
	const y = iota // y == 0
}

func shortIfStatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func elseIfStatement(a, b int) bool {
	if a < b {
		return true
	} else if a > b { // like pythons elif
		return false
	}
	return true
}

func switchStatement(i int) string {
	switch i {
	case 0:
		return "Zero"
	case 999:
		return "999"
	default:
		return "Default value"
	}
}

func fallthroughStatement(i int) string {
	switch i {
	case 999:
		fallthrough // the next case will be execute
		//anyway, no matter false or true statement in next case
	case 1:
		return "pass fallthrough"
	default:
		return "Default value"
	}
}

func switchWithNoCondition(i int) int {
	switch { // no condition
	case i == 1:
		return i + 1
	default:
		return i
	}
}

func allNumbersDiff(i int) string {
	a := i / 100     // get first item
	b := i / 10 % 10 // get second item
	c := i % 10      // get third item
	switch {
	case a != b && a != c && b != c:
		return "YES"
	default:
		return "NO"
	}
}

func getFirstDigit(i int) int {
	switch {
	case i == 10000:
		return 1
	case i/1000 > 0:
		return i / 1000
	case i/1000 == 0 && i/100 > 0:
		return i / 100
	case i/1000 == 0 && i/100 == 0 && i/10 > 0:
		return i / 10
	default:
		return i
	}
}

func happyTicket(i int) string {
	first3 := i / 1000
	second3 := i % 1000
	switch {
	case helper(first3) == helper(second3):
		return "YES"
	default:
		return "NO"
	}
}

func helper(i int) int {
	a := i / 100
	b := i / 10 % 10
	c := i % 10
	return a + b + c
}

func leapYearCheck(i int) string {
	switch {
	case i%400 == 0:
		return "YES"
	case i%4 == 0 && i%100 != 0:
		return "YES"
	default:
		return "NO"
	}
}

func sumBetweenTwoNumber(a, b int) int {
	// var a, b int
	var sum int
	// fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func sumOfTwoInRange(s []int) int {
	var sum int

	for _, i := range s {
		if (i >= 10 && i <= 99) && i%8 == 0 {
			sum += i
		}
	}
	return sum
}

func countElements(s []int) int {
	var maxValue = 0
	var count = 1
	for _, current := range s {
		if current > maxValue {
			maxValue = current
		} else if current == maxValue {
			count++
		} else if current == 0 {
			break
		}
	}
	return count
}

func twoNumbersDigitCompare(x, y int) string {
	s := ""
	initial := 10000
	z := y

	for i := 0; i < 5; i++ { // first number
		currentCompareFirst := x / initial
		if currentCompareFirst > 0 { // starts with 5 digits number and so on
			x = x - currentCompareFirst*initial

			childInitial := 10000
			for j := 0; j < 5; j++ { // second number
				currentCompareSecond := y / childInitial
				if currentCompareSecond > 0 {
					y = y - currentCompareSecond*childInitial

					if currentCompareSecond == currentCompareFirst {
						s = s + fmt.Sprintf("%d ", currentCompareFirst)
					}
				}
				childInitial = childInitial / 10
			}

		}
		initial = initial / 10
		y = z
	}
	return s[:len(s)-1]
}

// FORMMATING
func cutFloat(x float64) string {
	if x <= 0 {
		return fmt.Sprintf("%2.2f", x)
	} else if x > 10000 {
		return fmt.Sprintf("%e", x)
	} else {
		return fmt.Sprintf("%.4f", x*x)
	}
}

//ARRAYS
func workWithArray(workArray [10]uint8, idxArray [6]uint8) [10]uint8 {
	for _, idx := range []int{0, 2, 4} {
		workArray[idxArray[idx]], workArray[idxArray[idx+1]] = workArray[idxArray[idx+1]], workArray[idxArray[idx]]
	}
	// for _, value := range workArray {
	// 	fmt.Printf("%v ", value)
	// }
	return workArray
}

func get4Element(sBig [5]int) int {
	return sBig[3]
}

func sumOfTree(i int) int {
	var a, b, c int
	a = i / 100
	b = i % 100 / 10
	c = i % 10
	return a + b + c
}

func reverseNumber(i int) string {
	var a, b, c int
	a = i / 100
	b = i % 100 / 10
	c = i % 10
	return fmt.Sprintf("%d%d%d", c, b, a)
}

func countHoursAndMinutes(i int) string {
	var hours, minutes int
	if i > 0 && i < 86399 {
		hours = i / 60 / 60
		minutes = (i - hours*60*60) / 60
	}
	return fmt.Sprintf("%d %d", hours, minutes)
}

func ifTriangleRight(a int, b int, c int) string {
	right := "Прямоугольный"
	notRight := "Непрямоугольный"
	switch {
	case a*a+b*b == c*c:
		return right
	case a*a+c*c == b*b:
		return right
	case b*b+c*c == a*a:
		return right
	default:
		return notRight
	}
}

func ifTriangleExist(a int, b int, c int) string {
	exist := "Существует"
	notExist := "Не существует"

	if a+b > c && a+c > b && b+c > a {
		return exist
	}
	return notExist
}

// func arithmeticMean(a int, b int) string {
// 	if
// }

func countZeros(s []int) int {
	var count int
	for _, i := range s {
		if i == 0 {
			count++
		}
	}
	return count
}

func countMinValue(s []int) int {
	var count, minValue int
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			break
		}
		if s[i+1] <= s[i] {
			minValue = s[i+1]
		}
	}
	for _, i := range s {
		if i == minValue {
			count++
		}
	}
	return count
}

func digitalNumberCode(x int) int {
	initial := 10000000
	s := 0

	for i := 0; i < 8; i++ {
		current := x / initial
		if current > 0 {
			x = x - current*initial
		}
		initial = initial / 10
		if current != 0 {
			s += current
		}
	}
	x = (s-s%10)/10 + (s - (s - s%10))
	return x
}

func calcDivisor(a int, b int) string {
	s := []int{}
	maxValueExist := false
	maxValue := 0

	for i := a; i != b+1; i++ {
		if i%7 == 0 {
			s = append(s, i)
		}
	}

	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			maxValueExist = true
			maxValue = s[i]
			break
		}
		if s[i+1] > s[i] {
			maxValueExist = true
			maxValue = s[i+1]
		}
	}

	if maxValueExist {
		return fmt.Sprintf("YES %d", maxValue)
	}
	return fmt.Sprintf("NO %d", maxValue)
}

func rightSuffix(i int) string {
	switch {
	case i%10 == 0 || (i >= 10 && i <= 20) || (i%10 == 5 && i >= 10) || (i > 20 && (i%10) >= 6 && (i%10) <= 9):
		return fmt.Sprintf("%d korov", i)
	case (i%10 == 2 || i%10 == 3 || i%10 == 4) && (i < 10 || i > 20):
		return fmt.Sprintf("%d korovy", i)
	case i%10 == 1:
		return fmt.Sprintf("%d korova", i)
	default:
		return fmt.Sprintf("%d korova", i)
	}
}

func powTo2(a int) string {
	x := 1
	res := ""
	res = res + fmt.Sprintf("%d", 1)
	for i := 1; i != a+1; i++ {
		for j := 1; j != i; j++ {
			x = 2 * x
			if x < a {
				res = res + fmt.Sprintf(" %d", x)
			} else {
				break
			}
		}
		if x > a {
			break
		}
	}
	return res

	// 2nd var
	// var n int
	// fmt.Scan(&n)

	// for i := 1; i <= n; i = i * 2 {
	//     fmt.Printf("%v ", i)
	// }
}

func removeDigitFromNumber(a int, b int) int {
	var result int
	shift := 1
	for {
		if a != 0 {
			if a%10 != b {
				result = result + (a%10)*shift
				shift = shift * 10
			}
			a = a / 10
		} else {
			break
		}
	}
	return result
}

func placeInFibonacciLine(a int) int {
	fib1 := 1
	fib2 := 1
	i := 3
	if a == 1 {
		return 1
	}
	for {
		fibSum := fib1 + fib2
		fib1 = fib2
		fib2 = fibSum

		if fib2 == a {
			break
		} else if fib2 > a {
			return -1
		}
		i++
	}
	return i
}
