package main

import (
	"fmt"

	"github.com/TafveezA/P2P-GoDApplication/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
