package main

import (
	"fmt"

	"github.com/TafveezA/P2P-GoDApplication/deck"
)

func main() {
	for j := 0; j < 10; j++ {
		d := deck.Shuffle(deck.New())
		fmt.Println(d)
		fmt.Println()
	}
}
