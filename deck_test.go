package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "1 of Spades" {
		t.Errorf("Expected 1 of spades, but got %v", d[0])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
