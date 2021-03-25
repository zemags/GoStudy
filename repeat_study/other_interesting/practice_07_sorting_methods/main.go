package main

func countingSort(a []int) []int {
	// k -max value for current array, should know or find

	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	c := make([]int, max+1)
	for _, v := range a {
		c[v]++
	}

	a = []int{}
	for i := 0; i < max+1; i++ {
		for j := 0; j < c[i]; j++ {
			a = append(a, i)
		}
	}

	return a
}

func selectionSort(a []int) []int {
	min := 0
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return []int{}
}

func main() {

}
