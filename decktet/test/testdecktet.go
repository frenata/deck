package main

import (
	"fmt"

	"github.com/frenata/gaga/decktet"
)

func main() {
	//ranks := decktet.BasicRanks
	//fmt.Println(ranks)
	//fmt.Println(decktet.Ace < ranks[1])
	//fmt.Println(decktet.Ace + decktet.Two)
	d := decktet.BasicDeck

	fmt.Printf("%v\n", d.Shuffled[0])
}
