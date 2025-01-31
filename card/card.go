package card

import (
	"fmt"
	"strings"
)

type Card struct {
	r Rank
	s Suit
}

func New(r Rank, s Suit) Card {
	return Card{r, s}
}

// String returns the string representation of a card
func (c Card) String() string {
	lines := []string{
		"┌───────┐",
		"│%-2s     │",
		"│       │",
		"│   %s   │",
		"│       │",
		"│     %-2s│",
		"└───────┘",
	}
	card_string_template := strings.Join(lines, "\n")
	card_string := fmt.Sprintf(card_string_template, c.r, c.s, c.r)
	return card_string
}
