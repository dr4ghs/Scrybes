package game

type Deck struct {
	Cards []Card
}

type Player struct {
	Deck   Deck
	Side   Deck
	Hand   []Card
	Damage uint
}

type Game struct {
	Player   Player
	Opponent Player
}
