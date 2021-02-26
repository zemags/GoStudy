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
	for _, s := range strings.Split(text, " ") {
		tempMap[s]++
	}

	values := make([]int, 0, len(tempMap))
	for _, v := range tempMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for _, v := range values {
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
	return res, nil
}

func pop(m map[string]int, key string) (int, bool) {
	v, ok := m[key]
	if ok {
		delete(m, key)
	}
	return v, ok
}
