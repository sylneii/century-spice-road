package game

import (
	"sync"

	"github.com/sylneii/century-spice-road/card"
)

type GameState struct {
	players       []player
	currentPlayer player
	round         int
	tradingDeck   tradingStack
	scoringDeck   scoringStack
	mu            sync.Mutex
}

type player struct {
	username     string
	caravan      caravan
	tradingCards []card.TradingCard
	scoringCards []card.ScoringCard
	goldCoins    int64
	silverCoins  int64
}

type caravan struct {
	spices map[card.Spice]int64
	limit  int64
}

type tradingStack struct {
	tradingCards []card.TradingCard
	spiceOnCards []map[card.Spice]int
}

type scoringStack struct {
	scoringCards  []card.ScoringCard
	goldCoins     int64
	silverCoins   int64
	isGoldTaken   bool
	isSilverTaken bool
}

func (c *caravan) IsFull() {

}
