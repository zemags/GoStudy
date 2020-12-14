package main

import (
	"fmt"
	"strings"
) // import multiply imports

// Create a new type of 'deck'
// which is a slice of string
type deck []string // new type deck extends existing type string, its like class declaration

func newDeck() deck { // return a value of type 'deck'
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newCard := value + " of " + suit
			cards = append(cards, newCard)
		}
	}

	return cards
}

// Create new function thats belongs to type deck
// any var of type 'deck' now gets access to the 'print' method
// d is cards from main.go, d it's like self in python
// by convention it  will be one or two/three letter to the reference type in our case 'deck'
func (d deck) print() { //func with name print now is recivier
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}

// (deck, deck) means return two values of type 'deck'
func deal(d deck, handSize int) (deck, deck) { // handSize - number of arguments, its argument inside of func call
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { // (d deck) means recivier func
	return strings.Join([]string(d), ",") // []string(d) return a slice of strings
}
