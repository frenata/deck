package deck

/* Card interface defines a object that represents a card in a card game
   The implementing struct should be able to represent the Card via the String()
   function.
   Cards can be collected and shuffled in a Deck and be held by Players.

   TODO: Should implementations store information about what Player has played a Card?
*/
type Card interface {
	//PlayedBy(Player) Player
	String() string
}

// PrintCards takes a slice of Cards and returns a single string with each Card printed
// and seperated by spaces.
// For use as a helper function for structs that implement Card,
func PrintCards(stack []Card) string {
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

// CardSlice takes a varidric list of Cards and returns a slice of Cards.
// Useful for taking a slice of CardType and returning a slice of Card.
func CardSlice(cards ...Card) []Card {
	slice := make([]Card, len(cards))

	for i, c := range cards {
		slice[i] = c
	}
	return slice
}

// PopCard attempts to remove a Card from a slice of Cards, returns true if successful.
// **Not currently being used**
func popCard(c Card, s []Card) bool {
	for i, v := range s {
		if c == v {
			s = append(s[:i], s[i+1:]...)
			return true
		}
	}
	return false
}
