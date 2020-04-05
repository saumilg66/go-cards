package main

import "fmt"

func main() {

	cards := newDeck()

	cards.saveToFile("myCurrentDeck.txt")

	hand, _ := deal(cards, 4)

	fmt.Println(hand.toString())
	fmt.Println("*******")
	//remainingHand.print()
}
