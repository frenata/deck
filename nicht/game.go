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
	center *NichtCard
	flip   []*NichtCard
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
	g.deck.Players = gp
	g.deck.DealAll()
	g.button <- g.players[0]
	g.log.Printf("Cards dealt, %v has the button.\n", g.players[0].Name)
	return g
}

// Prints the current gamestate
func (g *Game) String() string {
	s := "**Deck**\n"
	s += fmt.Sprintln(g.deck)
	s += "**Hands**\n"
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
	g.log.Printf("%v starts the round by putting %v in the middle.\n", b.Name, g.board.center)
	for _, p := range g.players {
		if p != b {
			g.board.flip = append(g.board.flip, p.PlayRand())
		}
	}
	g.turn <- b
}

func (g *Game) printFlip() (s string) {
	for _, f := range g.board.flip {
		s += fmt.Sprintf("%v: %v | ", f.play, f)
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
		g.log.Printf("%v chooses %v\n", p.Name, g.board.flip[r])
		p.AddTable(g.board.flip[r])
		g.turn <- g.board.flip[r].play
		g.board.flip = append(g.board.flip[:r], g.board.flip[r+1:]...)
	} else { // center is the only card left
		p.AddTable(g.board.center)
		g.log.Printf("%v takes %v from the middle.\n", p.Name, g.board.center)
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
	var blueV, redV, yellowV, greenV int = 0, 0, 0, 0
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

	p.Score += total
	return score, total
}

// TODO: issues with this code or something related, cards not returning properly.
func (g *Game) Reshuffle() {
	for _, p := range g.players {
		g.deck.ReturnCards(p.Table)
		p.Table = nil
	}
	g.deck.Shuffle(-1)
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
	fmt.Println(g.Score())

	g.log.Println("Round 2")
	g.Reshuffle()
	g.deck.DealAll()
	for i := 0; i < 15; i++ {
		g.preRandRound()
		g.randRound()
	}
	fmt.Println(g.Score())

	g.log.Println("Round 3")
	g.Reshuffle()
	g.deck.DealAll()
	for i := 0; i < 15; i++ {
		g.preRandRound()
		g.randRound()
	}
	fmt.Println(g.Score())

	for _, p := range g.players {
		fmt.Printf("%v's final score: %v\n", p, p.Score)
	}

	ioutil.WriteFile("nicht.log", buf.Bytes(), 0660)
}
