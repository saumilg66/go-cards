package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingHand := deal(cards, 4)

	hand.print()
	fmt.Println("*******")
	remainingHand.print()
}
