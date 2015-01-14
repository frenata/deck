// nicht die bohn, initial gaga implementation
package gaga

import (
	"fmt"
	"math/rand"
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

var colors = [4]string{blue, green, yellow, red}
var vSpecial = [5]int{-1, -1, -1, 0, 2}
var vNormal = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

type Card struct {
	Color string
	Cat   string
	Value int
}

func NewCard(color, cat string, value int) *Card {
	c := new(Card)
	c.Color = color
	c.Cat = cat
	c.Value = value

	return c
}

func (c *Card) String() string {
	s := c.Color[:1]
	v := fmt.Sprint(c.Value)
	if c.Cat == special {
		return s + "x" + v
	} else { //if cat == normal {
		return s + v
	}
}

func PrintStack(stack []*Card) string {
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
	cards    []*Card
	shuffled []*Card
	dealt    []*Card
}

// could load from a config file
func NewDeck() *Deck {
	d := new(Deck)
	for _, color := range colors {
		for _, i := range vNormal {
			d.cards = append(d.cards, NewCard(color, normal, i))
		}
		for _, i := range vSpecial {
			d.cards = append(d.cards, NewCard(color, special, i))
		}
	}
	d.shuffled = make([]*Card, len(d.cards))
	d.dealt = make([]*Card, 0, len(d.cards))
	copy(d.shuffled, d.cards)
	return d
}

func (d *Deck) String() string {
	return PrintStack(d.shuffled)
}
func (d *Deck) PrintAll() string {
	return PrintStack(d.cards)
}
func (d *Deck) PrintDealt() string {
	return PrintStack(d.dealt)
}
func (p *Player) PrintHand() string {
	return PrintStack(p.Hand)
}
func (p *Player) PrintTable() string {
	return PrintStack(p.Table)
}

func (d *Deck) Shuffle() {
	rnd := newSeed()
	n := make([]*Card, len(d.shuffled), len(d.cards))
	r := rnd.Perm(len(d.shuffled))
	j := 0
	for _, i := range r {
		n[j] = d.shuffled[i]
		j++
	}
	d.shuffled = n
}

func (d *Deck) Deal(players []*Player) (n int) {
	for len(d.shuffled) > 0 {
		for _, p := range players {
			n++
			p.Hand = append(p.Hand, d.shuffled[0])
			d.dealt = append(d.dealt, d.shuffled[0])
			//fmt.Printf("Size is %v. Dealt %v to player %v\n", len(d.shuffled), d.shuffled[0], p.name)
			d.shuffled = d.shuffled[1:]
		}
	}
	return n
}

func newSeed() *rand.Rand {
	i63 := rand.Int63()
	s := rand.NewSource(i63)
	r := rand.New(s)
	return r
}

type Player struct {
	Name  string
	Hand  []*Card
	Table []*Card
}

func NewPlayer(name string) *Player {
	p := new(Player)
	p.Name = name
	return p
}

func (p *Player) PlayRand() *PlayedCard {
	rnd := newSeed()
	n := rnd.Intn(len(p.Hand))
	c := p.Hand[n]
	p.Hand = append(p.Hand[:n], p.Hand[n+1:]...)

	return &PlayedCard{c, p}
}

func (p *Player) AddTable(c *Card) {
	p.Table = append(p.Table, c)
}

type PlayedCard struct {
	Card   *Card
	Player *Player
}

func (pc *PlayedCard) String() string {
	return fmt.Sprintf("%v: %v", pc.Player.Name, pc.Card)
}

/*func main() {
	d := NewDeck()
	d.Shuffle()
	d.Shuffle()
	//fmt.Println(d)
	players := []*Player{
		NewPlayer("P1"),
		NewPlayer("P2"),
		NewPlayer("P3"),
		NewPlayer("P4"),
	}
	d.Deal(players)
	for _, p := range players {
		fmt.Println(p.name)
		fmt.Println(p.PrintHand())
	}

}*/
