package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// countWordsFreq()
}

func countWordsFreq(text string, freq int) (res []string, err error) {
	if text == "" {
		return res, fmt.Errorf("error: empty string provided")
	}

	tempMap := make(map[string]int)
	for _, s := range strings.Split(strings.Join(strings.Fields(text), " "), " ") {
		tempMap[s]++
	}

	values := make([]int, 0, len(tempMap))
	for _, v := range tempMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values[:freq])

	for _, v := range values[:freq] {
		tempSlice := []string{}
		for key, val := range tempMap {
			if val == v {
				tempSlice = append(tempSlice, key)
				delete(tempMap, key)
			}
		}
		sort.Strings(tempSlice)
		res = append(res, tempSlice...)
	}
	return res[:freq], nil
}
