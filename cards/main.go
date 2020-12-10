package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // : use when first initialize
	// card = "Five of Diamonds" // when second time reinitilize dont need :

	// card := newCard()
	// fmt.Println(card)

	cards := []string{newCard(), "New string"} // [] slice with type string
	fmt.Println(cards)

	cards = append(cards, "Appended string") //  append return new slice by appending value, NOT MODIFIED the older one

	for idx, card := range cards {
		fmt.Println(idx, card)
	}

}

func newCard() string { // return a value type string
	return "Six of Diamonds"
}
