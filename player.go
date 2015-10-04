package deck

// Player defines an object that can be given Cards, and can print the cards it holds
// as a String.
// TODO: String() should be removed or renamed to Hand() string or Cards() string,
// the same implementations kept, to better represent how it is being implemented.
type Player interface {
	AddCard(c Card)
	String() string
}
