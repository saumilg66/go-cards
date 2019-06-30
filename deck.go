package main

import "fmt"

type deck []string

// print all cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"1", "2", "3", "4"}

	for _, suite := range cardSuite {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

// deal cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
