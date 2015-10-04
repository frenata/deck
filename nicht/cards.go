package main

// TODO:
// DONE: seperate library and game logic, very messy atm
// DONE: fix rand seed setup so I can optionally generate new shuffles/games each runtime.
// DONE: log game actions to file, for checking of game logic
// DONE: lots of things should be interfaces, for further game dev: Shuffler, Player, Scorer?
// implement naive AI, instead of taking random cards, take the card that increases personal score the most
// much further on: implement 'smart' AI, take the card that increases score delta by the most
// (vs. currently highest scoring player?)

import (
	"fmt"
	"math/rand"

	"github.com/frenata/deck"
)

// defining some string constants for use in defining the standard nicht deck.
const (
	blue   = "Blue"
	green  = "Green"
	yellow = "Yellow"
	red    = "Red"

	special = "special"
	normal  = "normal"
)

// the types of card *details* contained in the nicht deck: color, special numbers, regular numbers.
var colors = [4]string{blue, green, yellow, red}
var vSpecial = [5]int{-1, -1, -1, 0, 2}
var vNormal = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// some barely named AI players. This might make more sense in game.go
var nichtPlayers = []*NichtPlayer{
	NewNichtPlayer("P1"),
	NewNichtPlayer("P2"),
	NewNichtPlayer("P3"),
	NewNichtPlayer("P4"),
}

// NewNichtDeck creates and populates a standard Nicht Deck
// TODO: Since this really only needs to be done once, move the logic to an init()
// function and create a package variable that holds a standard Nicht Deck.
func NewNichtDeck() *deck.Deck {
	c := make([]deck.Card, 60)
	n := 0
	for _, color := range colors {
		for _, i := range vNormal {
			c[n] = NewNichtCard(color, normal, i)
			n++
		}
		for _, i := range vSpecial {
			c[n] = NewNichtCard(color, special, i)
			n++
		}
	}
	return deck.New(c)
}

// A NichtCard contains the vital information about a Nicht Card, color, value,
// whether it is special or normal, and who played it.
type NichtCard struct {
	Color string
	Cat   string
	Value int
	play  *NichtPlayer
}

// NewNichtCard creates a new card.
// TODO: Actually, this seems like it can be easily removed in favor of a struct literal in NewNichtDeck?
// TODO: Not needed by the user, does not need to be exported.
func NewNichtCard(color, cat string, value int) *NichtCard {
	c := new(NichtCard)
	c.Color = color
	c.Cat = cat
	c.Value = value

	return c
}

// String prints a NichtCard representation:
// 	first letter of the color
// 	'x' if the card is special
//	the integer value of the card
func (c *NichtCard) String() string {
	s := c.Color[:1]
	v := fmt.Sprint(c.Value)
	if c.Cat == special {
		return s + "x" + v
	} else { //if cat == normal {
		return s + v
	}
}

// On further thought, this abstraction isn't producing any benefit.
// NichtCard is merely fulfilling the interface mindlessly,
// the actual game logic in game.go is interacting with the struct
// directly.
// Having the Card interface inherently related to a Player might
// be worthwhile, but only if deck.Deck also holds a []Player, so that
// by itself deck.Deck could handle removing Player pointers to cards
// and reshuffling those cards into the deck. That seems a small benefit
// for an unwieldy extra abstraction.
/*
func (c *NichtCard) PlayedBy(p deck.Player) deck.Player {
	if p == nil {
		c.play = nil
	} else {
		c.play = p.(*NichtPlayer)
	}
	return c.play
}
*/

// A NichtPlayer is a player of the Nicht game, has a name, a score, a hand of cards, and a table of cards.
type NichtPlayer struct {
	Name  string
	Hand  []deck.Card
	Table []deck.Card
	Score int
}

// NewNichtPlayer creates a new player with their name.
func NewNichtPlayer(name string) *NichtPlayer {
	p := new(NichtPlayer)
	p.Name = name
	return p
}

// Initial mock of AI play to make the game logic work, this simply randomly chooses an available card
// to play from hand.
// TODO: Obviously there should be a more intelligent function at some point.
func (p *NichtPlayer) PlayRand() *NichtCard {
	n := rand.Intn(len(p.Hand))
	c := p.Hand[n].(*NichtCard)
	p.Hand = append(p.Hand[:n], p.Hand[n+1:]...)
	c.play = p

	return c
}

// AddTable adds a card to a player's table.
func (p *NichtPlayer) AddTable(c *NichtCard) {
	p.Table = append(p.Table, c)
}

// AddCard adds a card to a player's hand. Implements deck.Player
func (p *NichtPlayer) AddCard(c deck.Card) {
	p.Hand = append(p.Hand, c.(*NichtCard))
}

// PrintHand prints a player's current hand.
func (p *NichtPlayer) PrintHand() string {
	var gc []deck.Card
	for _, c := range p.Hand {
		gc = append(gc, c)
	}
	return deck.PrintCards(gc)
}

// PrintTable prints a player's current table.
func (p *NichtPlayer) PrintTable() string {
	var gc []deck.Card
	for _, c := range p.Table {
		gc = append(gc, c)
	}
	return deck.PrintCards(gc)
}

// String prints a player's name. Should this also print Hand and Table?
func (p *NichtPlayer) String() string {
	return p.Name
}
