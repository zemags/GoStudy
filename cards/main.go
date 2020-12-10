package main

func main() {
	// cards := []string{newCard(), "New string"} // [] slice with type string

	cards := deck{newCard(), "Ace of Spades"} // initialize by deck type from deck.go it'l like instance of python class

	cards = append(cards, "Ace of Diamonds") //  append return new slice by appending value, NOT MODIFIED the older one

	cards.print()

}

func newCard() string { // return a value type string
	return "Six of Diamonds"
}
