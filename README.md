Go library for card game development

####Gaga (Deck Library)
* DONE: interfaces for Card and Player
* DONE: Deck struct, handles shuffling of Cards and dealing to Players
* TODO: other common interfaces of use? Scorer? more abstracter Player/CardHolder? Game?
* TODO: add Play(f func) method to Player, abstract out play strategy from clients?

####Nicht
* DONE: seperate library and game logic, very messy atm
* DONE: fix rand seed setup so I can optionally generate new shuffles/games each runtime.
* DONE: log game actions to file, for checking of game logic
* PARTIAL: lots of things should be interfaces, for further game dev: Shuffler, Player, Scorer?
* TODO: implement naive AI, instead of taking random cards, take the card that increases personal score the most, much further on: implement 'smart' AI, take the card that increases score delta by the most (vs. currently highest scoring player?)

####Decktet
* DONE: Basic deck definition
* TODO: Expanded deck: excuse, pawns, courts.

####Adaman
* DONE: deal cards correctly and print them in appropriate rows
* TODO: functions for evaluating whether a card can be claimed with a given set
* TODO: AI function to check all possible claims, find the most efficient ones
* TODO: program to run the Adaman-AI program, count statistics
