package main

import (
	"fmt"
	"sort"

	"github.com/frenata/deck"
	"github.com/frenata/deck/deck52"
)

type heartsPlayer struct {
	hand   []deck.Card
	tricks [][]deck.Card
	name   string
	score  int
}

func NewPlayer(name string) *heartsPlayer {
	p := &heartsPlayer{name: name}
	p.hand = make([]deck.Card, 0, 13)
	p.score = 0
	return p
}

func (p *heartsPlayer) String() string {
	sort.Sort(deck52.BySuit(p.hand))
	return fmt.Sprintf("%s: \n %s\nTricks: %d Score: %d\n", p.name, p.hand, len(p.tricks), p.score)
}

func (p *heartsPlayer) AddCard(c deck.Card) {
	p.hand = append(p.hand, c)
}
