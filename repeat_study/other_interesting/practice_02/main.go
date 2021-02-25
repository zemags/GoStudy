package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`(?i)[\d\+]+`)

func main() {
	// countWordsFreq()
}

func countWordsFreq(text string, freq int) (res []string, err error) {
	if text == "" {
		return res, fmt.Errorf("error: empty string provided")
	}

	tempMap := make(map[string]int)
	for _, s := range strings.Split(text, " ") {
		if _, inMap := tempMap[s]; !inMap {
			tempMap[s] = 1
		} else {
			tempMap[s] = tempMap[s] + 1
		}
	}

	values := make([]int, 0, len(tempMap))
	for _, v := range tempMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// tempSlice := []string{}
	// for _, v := range values {

	// }

	fmt.Println(tempMap)
	return res, nil
}

func pop(m map[string]int, key string) (int, bool) {
	v, ok := m[key]
	if ok {
		delete(m, key)
	}
	return v, ok
}
