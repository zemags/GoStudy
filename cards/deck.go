package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // writing to file, 0666-permissions anyone can read\write
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // bs - byteslice, err - error object\value of type error can return error or null\nil
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // 1 or any number means error, 0 - means success
	}

	s := strings.Split(string(bs), ",") // s - slice of strings name is convention
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano()) // time.Now().UnixNano() - return int64 time, its seed to generate
	r := rand.New(source)                           // its type of Rand

	for idx := range d {
		newPosition := r.Intn(len(d) - 1)

		d[idx], d[newPosition] = d[newPosition], d[idx]

	}
}
