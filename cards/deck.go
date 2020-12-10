package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck []string // new type deck extends existing type string, its like class declaration

// Create new function thats belongs to typ deck
// any var of type 'deck' now gets access to the 'print' method
// d is cards from main.go, d it's like self in python
// by convention it  will be one or two/three letter to the reference type in our case 'deck'
func (d deck) print() { //func with name print now is recivier
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}
