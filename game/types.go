package game

import (
	"sync"

	"github.com/sylneii/century-spice-road/card"
)

type GameState struct {
	players            []*player
	currentPlayerIndex int64
	round              int64
	tradingDeck        tradingStack
	scoringDeck        scoringStack
	pendingAction      pendingAction
	mu                 sync.Mutex
}

type player struct {
	username     string
	caravan      caravan
	currentCard  card.TradingCard
	usedCards    map[string]card.TradingCard
	tradingCards map[string]card.TradingCard
	scoringCards []card.ScoringCard
	goldCoins    int64
	silverCoins  int64
}

type caravan struct {
	spices map[card.Spice]int64
}

const (
	caravanLimit int64 = 10
)

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

type pendingAction interface {
	isPending()
}

type pendingActionDiscard struct {
}

type pendingActionUpgrade struct {
	remainingLevels int64
}

type pendingActionExchange struct {
	upgradesDone int64
}

func (p pendingActionDiscard) isPending()  {}
func (p pendingActionUpgrade) isPending()  {}
func (p pendingActionExchange) isPending() {}
