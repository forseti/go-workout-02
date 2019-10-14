package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_NewDeck(t *testing.T) {
	d := newDeck()

	fmt.Println("len: ", len(d))

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected Ace of Diamonds, but got %v", d[0])
	}

	lastIndex := len(d) - 1

	if d[lastIndex] != "Three of Clubs" {
		t.Errorf("Expected Three of Clubs, but got %v", d[lastIndex])
	}
}

func Test_SaveToNewDeck_And_NewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
