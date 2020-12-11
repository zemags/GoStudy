package main

func main() {
	// cards := []string{newCard(), "New string"} // [] slice with type string

	cards := newDeck()

	// cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
