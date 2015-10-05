// Hearts implementation
package main

import (
	"fmt"
	"math/rand"

	"github.com/frenata/deck"
	"github.com/frenata/deck/deck52"
)

func main() {
	fmt.Println("Play Hearts!")

	d := deck52.New()

	d.Shuffle()

	players := []deck.Player{NewPlayer("Andrew"), NewPlayer("Bekah"), NewPlayer("Chuck"), NewPlayer("Dave")}

	d.DealAll(players)

	fmt.Println(players)
	round := 1
	winner := playRandRound(players, 0)
	fmt.Printf("%s wins the round!\n", players[winner].(*heartsPlayer).name)
	fmt.Println(players)
	for round < 13 {
		round++
		fmt.Printf("Round: %d\n", round)
		winner = playRandRound(players, winner)
		fmt.Printf("%s wins the round!\n", players[winner].(*heartsPlayer).name)
		fmt.Println(players)
	}
}

func playRand(p deck.Player, i int, trick []deck.Card, lead deck.Card) deck.Card {
	hp := p.(*heartsPlayer)
	var c deck.Card

	if lead != nil {
		hl := lead.(deck52.Card)
		//fmt.Println(hl)
		if hasSuit(hp, hl.Suit()) {
			//fmt.Println("Can follow suit!")
			c = randHand(hp, hl.Suit())
			trick[i] = c
			fmt.Printf("%s follows suit and plays %s\n", hp.name, c)
			return c
		}
	}

	c = randHand(hp, "")
	trick[i] = c
	if lead != nil {
		fmt.Printf("%s does not follow suit and plays %s\n", hp.name, c)
	}
	return c
}

func playRandRound(players []deck.Player, leader int) (winner int) {
	trick := make([]deck.Card, len(players))
	i := leader

	lead := playRand(players[i], i, trick, nil)
	fmt.Printf("%s leads with %s\n", players[i].(*heartsPlayer).name, lead)
	//fmt.Println(lead)
	for j := 0; j < 3; j++ {
		i++
		if i == 4 {
			i = 0
		}
		playRand(players[i], i, trick, lead)
	}
	//fmt.Println(trick)

	max := lead.(deck52.Card)
	winner = leader
	for i := 0; i < 4; i++ {
		test := trick[i].(deck52.Card)
		if b, err := max.Less(test); err == nil && b {
			//fmt.Printf("New best card: %s\n", test)
			max = test
			winner = i
		}
	}

	wp := players[winner].(*heartsPlayer)
	wp.tricks = append(wp.tricks, trick)
	return winner
}

func randHand(p *heartsPlayer, s deck52.Suit) (c deck.Card) {
	if s == "" {
		i := rand.Intn(len(p.hand))
		c = p.hand[i]
		//p.hand = append(p.hand[:i], p.hand[i+1:]...)
	} else {
		match := filterSuit(p.hand, s)
		i := rand.Intn(len(match)) //need to delete item
		c = match[i]
	}

	removeCard(p, c)
	return c
}

func hasSuit(p *heartsPlayer, s deck52.Suit) bool {
	for _, c := range p.hand {
		if c.(deck52.Card).Suit() == s {
			return true
		}
	}
	return false
}

func filterSuit(cards []deck.Card, s deck52.Suit) (results []deck.Card) {
	for _, c := range cards {
		if c.(deck52.Card).Suit() == s {
			results = append(results, c)
		}
	}
	return results
}

func removeCard(p *heartsPlayer, c deck.Card) bool {
	for i, v := range p.hand {
		if v == c {
			p.hand = append(p.hand[:i], p.hand[i+1:]...)
			return true
		}
	}
	return false
}
