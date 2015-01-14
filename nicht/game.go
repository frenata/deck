package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/frenata/gaga"
)

type Game struct {
	players []*NichtPlayer
	turn    chan *NichtPlayer
	button  chan *NichtPlayer
	deck    *gaga.Deck
	board   *Board
	log     *log.Logger
}

type Board struct {
	center *PlayedCard
	flip   []*PlayedCard
}

func NewNichtGame(players []*NichtPlayer, l *log.Logger) *Game {
	g := new(Game)
	g.players = players
	g.deck = NewNichtDeck()
	g.board = &Board{}
	g.button = make(chan *NichtPlayer, 1)
	g.turn = make(chan *NichtPlayer, 1)
	g.log = l

	g.deck.Shuffle(1)
	g.log.Println("deck shuffled")
	var gp []gaga.Player
	for _, p := range g.players {
		gp = append(gp, p)
	}
	g.deck.DealAll(gp)
	g.button <- g.players[0]
	g.log.Printf("Cards dealt, %v has the button.\n", g.players[0].Name)
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
	g.log.Printf("%v starts the round by putting %v in the middle.\n", b.Name, g.board.center.C.Show())
	for _, p := range g.players {
		if p != b {
			g.board.flip = append(g.board.flip, p.PlayRand())
		}
	}
	g.turn <- b
}

func (g *Game) printFlip() (s string) {
	for _, f := range g.board.flip {
		s += fmt.Sprintf("%v: %v | ", f.P.String(), f.C.Show())
	}
	return s
}

func (g *Game) randRound() {
	g.log.Printf("Cards available: %v\n", g.printFlip())
	g.randTurn(<-g.turn)
	g.randTurn(<-g.turn)
	g.randTurn(<-g.turn)
	g.randTurn(<-g.turn)
}

func (g *Game) randTurn(p *NichtPlayer) {
	if len(g.board.flip) > 0 {
		r := rand.Intn(len(g.board.flip))
		g.log.Printf("%v chooses %v\n", p.Name, g.board.flip[r].C.Show())
		p.AddTable(g.board.flip[r].C.(*NichtCard))
		g.turn <- g.board.flip[r].P.(*NichtPlayer)
		g.board.flip = append(g.board.flip[:r], g.board.flip[r+1:]...)
	} else { // center is the only card left
		p.AddTable(g.board.center.C.(*NichtCard))
		g.log.Printf("%v takes %v from the middle.\n", p.Name, g.board.center.C.Show())
		g.button <- p
		g.board.center = nil
		g.board.flip = nil
	}
}

func (g *Game) Score() (score string) {
	score = "**Scores**\n"
	for _, p := range g.players {
		ps, _ := PlayerScore(p)
		score += ps
	}
	return score
}

func PlayerScore(p *NichtPlayer) (score string, total int) {
	var blueV, redV, yellowV, greenV int
	var blueS, redS, yellowS, greenS int = 1, 1, 1, 1

	for _, gc := range p.Table {
		c := gc.(*NichtCard)
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
	var buf bytes.Buffer
	l := log.New(&buf, "Nicht: ", log.Ltime)
	g := NewNichtGame(nichtPlayers, l)
	g.log.Println("New Game!")
	g.deck.Shuffle(1)

	for i := 0; i < 15; i++ {
		g.preRandRound()
		g.randRound()
	}

	fmt.Println(g)
	fmt.Println(g.Score())
	ioutil.WriteFile("nicht.log", buf.Bytes(), 0660)
}
