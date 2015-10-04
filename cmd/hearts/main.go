// Hearts implementation
package main

import (
	"fmt"
	"sort"

	"github.com/frenata/deck/deck52"
)

func main() {
	fmt.Println("Play Hearts!")

	d := deck52.New()
	c1 := d.Cards()[51]

	fmt.Println(c1)
	fmt.Println(c1.(deck52.Card).Name())

	d.Shuffle(1)
	cards := d.Cards()
	sort.Sort(deck52.BySuit(cards))
	fmt.Println(cards)
}
