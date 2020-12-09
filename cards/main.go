package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // : use when first initialize
	// card = "Five of Diamonds" // when second time reinitilize dont need :

	card := newCard()
	fmt.Println(card)

	number := newNumber()
	fmt.Println(number)
}

func newCard() string { // return a value type string
	return "Six of Diamonds"
}

func newNumber() int {
	return 999
}
