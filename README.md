Go library for card game development, focusing on custom and non-standard cards and decks.

Package deck focuses on implementation of a Deck struct that can shuffle, deal cards, etc, along with interface definitions Card and Player for clients to use with the Deck Struct.

cmd/* contains example implementations
- the simple card game Nicht Die Bonne, using a non-standard deck of cards.
- TODO: standard 52 card deck (in the deck package) and a simple game or two using it.
