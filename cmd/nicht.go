package main

import (
	"fmt"
	"math/rand"
)
import "gaga"

var nichtPlayers = []*gaga.Player{
	gaga.NewPlayer("P1"),
	gaga.NewPlayer("P2"),
	gaga.NewPlayer("P3"),
	gaga.NewPlayer("P4"),
}

type Game struct {
	players []*gaga.Player
	turn    chan *gaga.Player
	button  chan *gaga.Player
	deck    *gaga.Deck
	board   *Board
}

type Board struct {
	center *gaga.PlayedCard
	flip   []*gaga.PlayedCard
}

func NewGame(players []*gaga.Player) *Game {
	g := new(Game)
	g.players = players
	g.deck = gaga.NewDeck()
	g.board = &Board{}
	g.button = make(chan *gaga.Player, 1)
	g.turn = make(chan *gaga.Player, 1)

	g.deck.Shuffle()
	g.deck.Deal(g.players)
	g.button <- g.players[0]
	return g
}

func (g *Game) String() string {
	s := "**Hands**\n"
	for _, p := range g.players {
		s += fmt.Sprintln(p.Name)
		s += fmt.Sprintln(p.PrintHand())
	}
	s += "**Scores**\n"
	for _, p := range g.players {
		s += fmt.Sprintln(p.Name)
		s += fmt.Sprintln(p.PrintTable())
	}
	s += "**Board**\n"
	if g.board.center != nil {
		s += "Center:\n" + g.board.center.String() + "\n"
	}
	if g.board.flip != nil {
		s += "Flip:\n"
		for _, f := range g.board.flip {
			s += f.String() + "\n"
		}
	}
	return s
}

func (g *Game) preRandRound() {
	b := <-g.button
	g.board.center = b.PlayRand()
	for _, p := range g.players {
		if p != b {
			g.board.flip = append(g.board.flip, p.PlayRand())
		}
	}
	g.turn <- b
}

func (g *Game) randRound() {
	g.randTurn(<-g.turn)
	g.randTurn(<-g.turn)
	g.randTurn(<-g.turn)
	g.randTurn(<-g.turn)
}

func (g *Game) randTurn(p *gaga.Player) {
	if len(g.board.flip) > 0 {
		r := rand.Intn(len(g.board.flip))
		p.AddTable(g.board.flip[r].Card)
		g.turn <- g.board.flip[r].Player
		g.board.flip = append(g.board.flip[:r], g.board.flip[r+1:]...)
	} else { // center is the only card left
		p.AddTable(g.board.center.Card)
		g.button <- g.board.center.Player
		g.board.center = nil
		g.board.flip = nil
	}
}

func main() {
	g := NewGame(nichtPlayers)
	fmt.Println(g)
	for i := 0; i < 15; i++ {
		g.preRandRound()
		g.randRound()
	}

	fmt.Println(g)
}
