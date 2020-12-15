package main

func main() {
	checkNumbers()
}

func checkNumbers() []string {
	sliceOfInts := []int{0,1,2,3,4,5,6,7,8,9,10}
	var results []string
	currentResult := "even"
	for _, number := range sliceOfInts {
		if number % 2 == 0 {
			currentResult = "even"		
		} else {
			currentResult = "odd"	
		}
		results = append(results, currentResult)
	}
	return results
}