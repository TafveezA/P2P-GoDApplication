package deck

import "fmt"

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Hearts:
		return "HEARTS"
	case Clubs:
		return "CLUBS"
	case Diamonds:
		return "DIAMONDS"
	default:
		panic("Invalid card cuit")
	}

}

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

type Card struct {
	suit  Suit
	value int
}

func (c Card) String() string {
	return fmt.Sprintf("%d of %s %s", c.value, c.suit)
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of the card cannot be higher than 13")
	}
	return Card{
		suit:  s,
		value: v,
	}
}

func suitToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	default:
		panic("Invalid card cuit")
	}
}
