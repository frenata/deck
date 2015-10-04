package deck52

import (
	"fmt"

	"github.com/frenata/deck"
)

// BySuit sorts cards first by suit (in Bridge order), then by rank, Ace high.
type BySuit []deck.Card

func (b BySuit) Len() int      { return len(b) }
func (b BySuit) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b BySuit) Less(i, j int) bool {
	bi := b[i].(Card)
	bj := b[j].(Card)
	switch {
	case bi.suit == bj.suit:
		r, _ := bi.Less(bj)
		return r
	case bi.suit == Clubs:
		return true
	case bi.suit == Diamonds && bj.suit == Hearts:
		return true
	case bi.suit == Diamonds && bj.suit == Spades:
		return true
	case bi.suit == Hearts && bj.suit == Spades:
		return true
	default:
		return false
	}
}

// Less evalutes whether a card is less than another card.
// If cards are not the same suit, returns error.
// Ace is high in this implelmentation
func (c Card) Less(c2 Card) (bool, error) {
	switch {
	case c.suit != c2.suit:
		return false, fmt.Errorf("Suits are not the same, cards cannot be compared")
	case c.rank == c2.rank:
		return false, nil
	case c.rank == Ace:
		return false, nil
	case c2.rank == Ace:
		return true, nil
	default:
		return c.rank < c2.rank, nil
	}
}
