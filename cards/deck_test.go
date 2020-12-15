package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) { // 'Test' starts with capital 'T'
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %s", d[0])
	}
	
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {
	os.Remove("_decktesting") // remove all files witn name _decktesting

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}