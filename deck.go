// Gaga defines some useful interfaces and functions for manipulating card decks in games.
package gaga

import (
	"math"
	"math/rand"
	"time"
)

/* Card interface defines a object that represents a card in a card game
   The implementing struct should be able to represent the Card via the String()
   function.
   Cards can be collected and shuffled in a Deck and be held by Players.

   TODO: Should implementations store information about what Player has played a Card?
*/
type Card interface {
	//PlayedBy(Player) Player
	String() string
}

// Player defines an object that can be given Cards, and can print the cards it holds
// as a String.
// TODO: String() should be removed or renamed to Hand() string or Cards() string,
// the same implementations kept, to better represent how it is being implemented.
type Player interface {
	AddCard(c Card)
	String() string
}

// PrintCards takes a slice of Cards and returns a single string with each Card printed
// and seperated by spaces.
// For use as a helper function for structs that implement Card,
func PrintCards(stack []Card) string {
	var s string

	for _, c := range stack {
		if s != "" {
			s = s + " " + c.String()
		} else {
			s = c.String()
		}
	}
	return s
}

// Deck collects Cards into a usable collection, suitable for shuffling and dealing.
// Three arrays are available:
// 	Cards - all the cards assigned to the deck, regardless of whether they have
//		been shuffled or dealt.
//	Shuffled - all the cards that have been shuffled but not yet dealt to Players.
//	Dealt - all the cards that have been dealt to Players.
// TODO: Should these arrays be accessible via getter functions only? So users cannot
// disrupt the data structure by removing a Card from the Cards array but not another.
// TODO: Cards should either be removed or renamed to something more suitable to
// represent that it represents all Cards attached to the Deck regardless of state,
// and should not be used directly.
//	All []Card
// Likewise, Shuffled is currently
// confusing and does not accurately represent what it is, which is the working set of
// cards in the deck, ready to be dealt, whether shuffled or not.
// 	Cards []Card
// Add another slice:
//	Discard
type Deck struct {
	Cards    []Card
	Shuffled []Card
	Dealt    []Card
}

// NewDeck accepts a slice of Cards and creates a new Deck ready for use.
// Cards and Shuffled are both populated with equivalent slices.
// Shuffle must be called explicitly after a NewDeck is returned to shuffle the cards.
func NewDeck(cards []Card) *Deck {
	d := new(Deck)
	d.Cards = cards
	d.Shuffled = make([]Card, len(d.Cards))
	d.Dealt = make([]Card, 0, len(d.Cards))
	copy(d.Shuffled, d.Cards)
	return d
}

// String prints all the *shuffled* cards in the deck.
func (d *Deck) String() string {
	return PrintCards(d.Shuffled)
}

// Shuffle takes a seed and randomizes the cards contained in the Shuffled slice.
// TODO: refactor to remove seed argument. Allow user to set the seed once and directly,
// via a seperate method.
func (d *Deck) Shuffle(seed int) {
	rnd := deckSeed(seed)
	n := make([]Card, len(d.Shuffled), len(d.Cards))
	r := rnd.Perm(len(d.Shuffled))
	j := 0
	for _, i := range r {
		n[j] = d.Shuffled[i]
		j++
	}
	d.Shuffled = n
}

// ReturnCards takes a slice of cards (from a Player?) and adds them back into Shuffled.
// TODO: This may be better represented by a seperate slice, Discards?
// Likewise, since the implication is that these cards were previously dealt by
// this deck, there should be error checking to verify: if a card was not previously
// dealt, an error should be returned. And if it was, remove it from Dealt slice.
func (d *Deck) ReturnCards(cards []Card) {
	for _, c := range cards {
		d.Shuffled = append(d.Shuffled, c)
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
	if len(d.Shuffled) > 0 {
		p.AddCard(d.Shuffled[0])
		d.Dealt = append(d.Dealt, d.Shuffled[0])
		d.Shuffled = d.Shuffled[1:]
		return true
	} else {
		return false
	}
}

// Seed sets a random seed for Deck shuffling and dealing. Passing -1 uses the current
// time, anything else is static and suitable for testing, etc.
// TODO: Rewrite to provide user access and permanent struct seed.
func deckSeed(seed int) *rand.Rand {
	s64 := int64(seed)

	if s64 == -1 {
		s64 = time.Now().UnixNano()
	}
	s := rand.NewSource(s64)
	r := rand.New(s)
	return r
}

// PopCard attempts to remove a Card from a slice of Cards, returns true if successful.
// **Not currently being used**
func PopCard(c Card, s []Card) bool {
	for i, v := range s {
		if c == v {
			s = append(s[:i], s[i+1:]...)
			return true
		}
	}
	return false
}

// CardCombinations returns a list of all possible combinations of given a slice of Cards.
func CardCombinations(cards []Card) [][]Card {
	var results [][]Card
	set := Combination(len(cards))

	for _, s := range set {
		row := []Card{}
		for _, v := range s {
			row = append(row, cards[v])
		}
		results = append(results, row)
	}
	return results
}

// Combination returns a list of all possible integer combinations given an integer.
func Combination(n int) [][]int {
	var slice []int
	var results [][]int
	var b byte
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		b = byte(i)
		slice = []int{}
		for j := 0; j < n; j++ { //int(math.Pow(2, float64(n))); j++ {
			if b>>uint(j)&1 == 1 {
				slice = append(slice, j)
			}
		}
		if len(slice) != 0 {
			results = append(results, slice)
		}
	}
	return results
}
