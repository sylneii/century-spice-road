package main

import (
	"errors"
	"log/slog"
)

type spices struct {
	turmeric int
	saffron  int
	cardamom int
	clove    int
}

type scoringCard struct {
	points int
	spices
}

type scoringStack struct {
	scoringCards  []scoringCard
	goldCoins     int
	silverCoins   int
	isGoldTaken   bool
	isSilverTaken bool
}

func newScoringStack(scoringCards []scoringCard, goldCoins, silverCoins int, isGoldtaken, isSilverTaken bool) *scoringStack {
	return &scoringStack{
		scoringCards:  scoringCards,
		goldCoins:     goldCoins,   //default 2 * number of players
		silverCoins:   silverCoins, //default 2 * number of players
		isGoldTaken:   isGoldtaken,
		isSilverTaken: isSilverTaken,
	}
}

func (ss *scoringStack) acquire(cardIndex int) error {
	currentCard := ss.scoringCards[cardIndex]

	//invalid moves
	if cardIndex > len(ss.scoringCards)-1 {
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
