package deck52

import "github.com/frenata/deck"

var d *deck.Deck

// Prep a standard deck
func init() {
	cards := make([]deck.Card, 52)
	n := 0
	for r := Ace; r <= King; r++ {
		cards[n] = newCard(r, Hearts)
		cards[n+1] = newCard(r, Diamonds)
		cards[n+2] = newCard(r, Spades)
		cards[n+3] = newCard(r, Clubs)
		n += 4
	}

	d = deck.New(cards)
}

// New returns a copy of the standard 52-card deck.
func New() *deck.Deck {
	n := *d
	return &n
}
