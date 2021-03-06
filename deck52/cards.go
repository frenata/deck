// Package deck52 implements the standard 52 card deck for use with github.com/frenata/deck
package deck52

import (
	"fmt"
	"strconv"
)

type Suit string
type Rank int

const ( // Suits
	Hearts   Suit = "Hearts"
	Diamonds      = "Diamonds"
	Spades        = "Spades"
	Clubs         = "Clubs"
)

const (
	_        = iota
	Ace Rank = iota
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
	suit  Suit
	rank  Rank
	name  string
	short string
}

// create a new card, since this is the standard deck, this has no need to be exported
// standard deck will be created in init()
func newCard(r Rank, s Suit) Card {
	c := Card{suit: s, rank: r}
	c.name = fmt.Sprintf("%s of %s", longRank(r), s)
	c.short = fmt.Sprintf("%s%s", shortRank(r), string(s[0]))

	return c
}

func (c Card) Suit() Suit {
	return c.suit
}

// String returns the short name of a card. Ex: "KC" for "King of Clubs"
func (c Card) String() string {
	return c.short
}

// Name returns the long name of a card. Ex: "King of Clubs"
func (c Card) Name() string {
	return c.name
}

// convert Rank to short value
func shortRank(r Rank) string {
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

// convert Rank to full string for Name()
func longRank(r Rank) string {
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
