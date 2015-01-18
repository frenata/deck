// nicht die bohn, initial gaga implementation
package gaga

import (
	"math"
	"math/rand"
	"time"
)

type Card interface {
	//PlayedBy(Player) Player
	String() string
}

type Player interface {
	AddCard(c Card)
	String() string
}

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

type Deck struct {
	Cards    []Card
	Shuffled []Card
	Dealt    []Card
}

func NewDeck(cards []Card) *Deck {
	d := new(Deck)
	d.Cards = cards
	d.Shuffled = make([]Card, len(d.Cards))
	d.Dealt = make([]Card, 0, len(d.Cards))
	copy(d.Shuffled, d.Cards)
	return d
}

func (d *Deck) String() string {
	return PrintCards(d.Shuffled)
}

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

func (d *Deck) ReturnCards(cards []Card) {
	for _, c := range cards {
		d.Shuffled = append(d.Shuffled, c)
	}
}

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

func deckSeed(seed int) *rand.Rand {
	s64 := int64(seed)

	if s64 == -1 {
		s64 = time.Now().UnixNano()
	}
	s := rand.NewSource(s64)
	r := rand.New(s)
	return r
}

func PopCard(c Card, s []Card) bool {
	for i, v := range s {
		if c == v {
			s = append(s[:i], s[i+1:]...)
			return true
		}
	}
	return false
}

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
