package main

// TODO:
// seperate library and game logic, very messy atm
// fix rand seed setup so I can optionally generate new shuffles/games each runtime.
// log game actions to file, for checking of game logic
// lots of things should be interfaces, for further game dev: Shuffler, Player, Scorer?
// implement naive AI, instead of taking random cards, take the card that increases personal score the most
// much further on: implement 'smart' AI, take the card that increases score delta by the most
// (vs. currently highest scoring player?)

import (
	"fmt"
	"math/rand"

	"github.com/frenata/gaga"
)

const (
	blue   = "Blue"
	green  = "Green"
	yellow = "Yellow"
	red    = "Red"

	special = "special"
	normal  = "normal"

	double = 2
	zero   = 0
	neg    = -1
)

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
	s += "**Tables**\n"
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

func (g *Game) Score() (score string) {
	for _, p := range g.players {
		ps, _ := PlayerScore(p)
		score += ps
	}
	return score
}

func PlayerScore(p *gaga.Player) (score string, total int) {
	var blueV, redV, yellowV, greenV int
	var blueS, redS, yellowS, greenS int = 1, 1, 1, 1

	for _, c := range p.Table {
		if c.Cat == special {
			switch c.Color {
			case blue:
				blueS *= c.Value
			case red:
				redS *= c.Value
			case yellow:
				yellowS *= c.Value
			case green:
				greenS *= c.Value
			}
		} else { // c.Cat == normal
			switch c.Color {
			case blue:
				blueV += c.Value
			case red:
				redV += c.Value
			case yellow:
				yellowV += c.Value
			case green:
				greenV += c.Value

			}
		}
	}

	blueV *= blueS
	redV *= redS
	yellowV *= yellowS
	greenV *= greenS

	total = blueV + redV + yellowV + greenV

	score = fmt.Sprintf(
		"%v -- Blue: %v, Red: %v, Yellow: %v, Green: %v, Total: %v\n",
		p.Name, blueV, redV, yellowV, greenV, total)

	return score, total
}

func main() {
	g := NewGame(nichtPlayers)
	g.deck.Shuffle()
	for i := 0; i < 15; i++ {
		g.preRandRound()
		g.randRound()
	}

	fmt.Println(g)
	fmt.Println(g.Score())

}
