// Package deck defines some useful interfaces and functions for manipulating card decks in games.
package deck

import (
	"math/rand"
	"time"
)

// Deck collects Cards into a usable collection, suitable for shuffling and dealing.
//	cards - all the cards that have been shuffled but not yet dealt to Players.
//	discards - all the cards that have been returned to the deck
type Deck struct {
	cards    []Card //cards currently in the deck
	discards []Card //cards currently in the deck
	rng      *rand.Rand
}

// New accepts a slice of Cards and creates a new Deck ready for use.
// Cards and Shuffled are both populated with equivalent slices.
// Shuffle must be called explicitly after a NewDeck is returned to shuffle the cards.
func New(cards []Card) *Deck {
	d := new(Deck)
	d.cards = make([]Card, len(cards))
	copy(d.cards, cards)
	d.discards = make([]Card, 0, len(d.cards))

	d.Seed(1)
	return d
}

// Cards returns a list of the cards in the deck.
func (d *Deck) Cards() []Card {
	cards := make([]Card, len(d.cards))
	copy(cards, d.cards)
	return cards
}

// Discards returns a list of the cards in the deck.
func (d *Deck) Discards() []Card {
	discards := make([]Card, len(d.discards))
	copy(discards, d.discards)
	return discards
}

// String prints all the *shuffled* cards in the deck.
func (d *Deck) String() string {
	return PrintCards(d.cards)
}

// Shuffle  shuffles the cards in the deck, based on the current seed.
func (d *Deck) Shuffle() {
	var toshuffle []Card
	toshuffle = append(toshuffle, d.cards...)
	toshuffle = append(toshuffle, d.discards...)

	d.discards = make([]Card, 0, len(toshuffle))

	shuffled := make([]Card, len(toshuffle))
	r := d.rng.Perm(len(toshuffle))
	j := 0
	for _, i := range r {
		shuffled[j] = toshuffle[i]
		j++
	}
	d.cards = shuffled
}

// ReturnCards takes a slice of cards (from a Player?) and adds them back into Shuffled.
// TODO: This may be better represented by a seperate slice, Discards?
// Likewise, since the implication is that these cards were previously dealt by
// this deck, there should be error checking to verify: if a card was not previously
// dealt, an error should be returned. And if it was, remove it from Dealt slice.
func (d *Deck) Discard(cards ...Card) {
	for _, c := range cards {
		d.discards = append(d.discards, c)
	}
}

// DealAll takes a slice of Players and deals cards to each in turn until
// none are left.
// This does *not* guarantee equal dealing.
// Returns number of cards dealt.
func (d *Deck) DealAll(players []Player) (n int) {
	for {
		for _, p := range players {
			if d.Deal(p) {
				n++
			} else { // if Deal fails, no more shuffled
				return n
			}
		}
	}
}

// Deal adds a card from the shuffled cards to a Player.
// Returns true if a card was dealt, false if not. (no cards left)
func (d *Deck) Deal(p Player) bool {
	if len(d.cards) > 0 {
		p.AddCard(d.cards[0])
		d.cards = d.cards[1:]
		return true
	} else {
		return false
	}
}

// Seed sets a random seed for Deck shuffling and dealing. Passing -1 uses the current
// time, anything else is static and suitable for testing, etc.
func (d *Deck) Seed(seed int) {
	s64 := int64(seed)

	if s64 == -1 {
		s64 = time.Now().UnixNano()
	}
	s := rand.NewSource(s64)
	d.rng = rand.New(s)
}
