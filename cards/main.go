package main

import "fmt"

func main() {
	// cards := []string{newCard(), "New string"} // [] slice with type string

	// cards := newDeck()

	// // cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// greeting := "Hello!"
	// fmt.Println([]byte(greeting))  // byteslice

	cards := newDeck()
	fmt.Println(cards.toString())
}
