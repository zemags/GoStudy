package main

import "fmt"

func main() {
	// maps a 'reference type', which means when we pass them to a function and modify them we are modifying the original map, not a copy
	// var colors map[string]string // empty map

	// colors := make(map[string]string)

	colors := map[string]string{ // in map keys must be one type, and value,must be one type, map like dict in python
		"red":   "#ff0000",
		"green": "#4bf745", // comma important!
	}

	colors["white"] = "#ffffff" // add value to map
	colors["undefined"] = "null"

	delete(colors, "undefined") // delete key-value pair

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key + " " + value)
	}
}
