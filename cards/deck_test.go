package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//test if the new deck has the right number of cards
	if len(d) != 52 {
		t.Errorf("Error:Expected deck length of 52, but got %v", len(d))
	}
}

func TestSaveToDckAndNwDeckFromFil(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.savetoFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {

		t.Errorf("Error:Expected deck length of 52, but got %v", len(loadedDeck))

	}
	os.Remove("_decktesting")
}
