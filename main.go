package main

import "fmt"

func main() {

	cards := newDeck()

	cards.shuffle()
	cards.print()
	//cards.saveToFile("myCurrentDeck.txt")

	//fromFile := readFromFile("myCurrentDeck.txt")

	//hand, _ := deal(cards, 4)

	//fmt.Println(hand.toString())
	fmt.Println("*******")
	//remainingHand.print()
}
