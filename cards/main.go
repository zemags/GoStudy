package main

func main() {
	// cards := []string{newCard(), "New string"} // [] slice with type string

	// cards := newDeck()

	// // cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// greeting := "Hello!"
	// fmt.Println([]byte(greeting))  // byteslice

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
