package main

import (
	"errors"
	"log/slog"

	"github.com/sylneii/century-spice-road/card"
)

type spice int64

const (
	turmeric spice = iota + 1
	saffron
	cardamom
	cinnamon
)

type scoringStack struct {
	scoringCards  []card.ScoringCard
	goldCoins     int64
	silverCoins   int64
	isGoldTaken   bool
	isSilverTaken bool
}

func newScoringStack(scoringCards []scoringCard, goldCoins, silverCoins int64, isGoldtaken, isSilverTaken bool) *scoringStack {
	return &scoringStack{
		scoringCards:  scoringCards,
		goldCoins:     goldCoins,   //default 2 * number of players
		silverCoins:   silverCoins, //default 2 * number of players
		isGoldTaken:   isGoldtaken,
		isSilverTaken: isSilverTaken,
	}
}

func (ss *scoringStack) acquire(cardIndex int64) error {
	currentCard := ss.scoringCards[cardIndex]

	//invalid moves
	if cardIndex > int64(len(ss.scoringCards))-1 {
		slog.Error("total_cards_less_than_index")
		return errors.New("Invalid move, pick a card on the display")
	}
	if cardIndex > 4 {
		slog.Error("scoring_card_not_on_display")
		return errors.New("Invalid move, pick a card on the display")
	}

	//acquire card
	ss.scoringCards = append(ss.scoringCards[:cardIndex], ss.scoringCards[cardIndex+1:]...) //use slices.Delete later
	slog.Info("player acquired scoring card", "scoring card", currentCard)

	// acquire gold or silver coin
	if cardIndex == 1 {
		if !ss.isGoldTaken {
			ss.goldCoins -= 1
			slog.Info("player acquired gold coin")
		}
		if ss.isGoldTaken && !ss.isSilverTaken {
			ss.silverCoins -= 1
			slog.Info("player acquired silver coin")
		}
	}
	if cardIndex == 2 {
		if !ss.isGoldTaken && !ss.isSilverTaken {
			ss.silverCoins -= 1
			slog.Info("player acquired silver coin")
		}
	}

	return nil
}

type tradingStack struct {
	tradingCards []tradingCard
	spiceOnCards []map[spice]int
}

func newTradingStack(tradingCards []tradingCard) *tradingStack {
	return &tradingStack{
		tradingCards: tradingCards,
		spiceOnCards: []map[spice]int{},
	}
}

func (ts *tradingStack) acquire(cardIndex int, putDownSpices []spice) map[spice]int {
	if cardIndex > 5 {
		slog.Error("invalid move")
	}
	acquiredSpices := ts.spiceOnCards[cardIndex]

	for idx := range ts.tradingCards[:cardIndex] {
		ts.spiceOnCards[idx][putDownSpices[idx]] += 1
	}
	ts.tradingCards = append(ts.tradingCards[:cardIndex], ts.tradingCards[cardIndex+1:]...)

	for idx := range ts.tradingCards[cardIndex:5] {
		ts.spiceOnCards[idx] = ts.spiceOnCards[idx+1] //check out of bounds for initialization
	}

	return acquiredSpices
}
