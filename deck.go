package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

// deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

// save to disk
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// read from File
func readFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\n")

	return deck(s)
}
