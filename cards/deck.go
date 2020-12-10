package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck []string

// Create new function thats belongs to typ deck
func (d deck) print() { //func without name is recivier
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}
