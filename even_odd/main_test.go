package main

import "testing"

func TestCheckNumbers(t *testing.T) {
	res := checkNumbers()
	trueRes := []string{"even", "odd", "even", "odd", "even", "odd", "even", "odd", "even", "odd", "even"}

	for idx := range trueRes {
		if trueRes[idx] != res[idx] {
			t.Errorf("Result doesnt match")
			break
		}
	}
}