// Package deck52 implements the standard 52 card deck for use with github.com/frenata/deck
package deck52

import (
	"fmt"
	"strconv"
)

type suit string
type rank int

const ( // suits
	Hearts   suit = "Hearts"
	Diamonds      = "Diamonds"
	Spades        = "Spades"
	Clubs         = "Clubs"
)

const (
	_        = iota
	Ace rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// a card for a standard 52-card deck
type Card struct {
	suit
	rank
	name  string
	short string
}

// create a new card, since this is the standard deck, this has no need to be exported
// standard deck will be created in init()
func newCard(r rank, s suit) Card {
	c := Card{suit: s, rank: r}
	c.name = fmt.Sprintf("%s of %s", longRank(r), s)
	c.short = fmt.Sprintf("%s%s", shortRank(r), string(s[0]))

	return c
}

// String returns the short name of a card. Ex: "KC" for "King of Clubs"
func (c Card) String() string {
	return c.short
}

// Name returns the long name of a card. Ex: "King of Clubs"
func (c Card) Name() string {
	return c.name
}

// convert rank to short value
func shortRank(r rank) string {
	switch r {
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return strconv.Itoa(int(r))
	}
}

// convert rank to full string for Name()
func longRank(r rank) string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return ""
	}
}
