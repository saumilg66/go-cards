package main

import "fmt"

func main() {

	cards := newDeck()

	cards.saveToFile("myCurrentDeck.txt")

	fromFile := readFromFile("myCurrentDeck.txt")

	hand, _ := deal(fromFile, 3)

	fmt.Println(hand.toString())
	fmt.Println("*******")
	//remainingHand.print()
}
