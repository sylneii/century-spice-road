package game

import (
	"sync"

	"github.com/sylneii/century-spice-road/cards"
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
	currentCard  cards.TradingCard
	usedCards    map[string]cards.TradingCard
	tradingCards map[string]cards.TradingCard
	scoringCards []cards.ScoringCard
	goldCoins    int64
	silverCoins  int64
}

type caravan struct {
	spices map[cards.Spice]int64
}

const (
	caravanLimit int64 = 10
)

type tradingStack struct {
	tradingCards []cards.TradingCard
	spiceOnCards []map[cards.Spice]int
}

type scoringStack struct {
	scoringCards  []cards.ScoringCard
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
