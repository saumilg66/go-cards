package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "1 of Spades" {
		t.Errorf("Expected 1 of spades, but got %v", d[0])
	}
}
