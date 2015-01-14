// nicht die bohn, initial gaga implementation
package gaga

import (
	"math/rand"
	"time"
)

type Card interface {
	Show() string
}

type Player interface {
	AddCard(c Card)
	String() string
}

func PrintCards(stack []Card) string {
	var s string

	for _, c := range stack {
		if s != "" {
			s = s + " " + c.Show()
		} else {
			s = c.Show()
		}

	}
	return s
}

type Deck struct {
	Cards    []Card
	Shuffled []Card
	Dealt    []Card
}

func (d *Deck) String() string {
	return PrintCards(d.Shuffled)
}

func (d *Deck) Shuffle(seed int) {
	rnd := NewSeed(seed)
	n := make([]Card, len(d.Shuffled), len(d.Cards))
	r := rnd.Perm(len(d.Shuffled))
	j := 0
	for _, i := range r {
		n[j] = d.Shuffled[i]
		j++
	}
	d.Shuffled = n
}

func (d *Deck) DealAll(players []Player) (n int) {
	for len(d.Shuffled) > 0 {
		for _, p := range players {
			n++
			p.AddCard(d.Shuffled[0])
			d.Dealt = append(d.Dealt, d.Shuffled[0])
			//fmt.Printf("Size is %v. Dealt %v to player %v\n", len(d.shuffled), d.shuffled[0], p.name)
			d.Shuffled = d.Shuffled[1:]
		}
	}
	return n
}

func NewSeed(seed int) *rand.Rand {
	s64 := int64(seed)

	if s64 == -1 {
		s64 = time.Now().UnixNano()
	}
	s := rand.NewSource(s64)
	r := rand.New(s)
	return r
}
