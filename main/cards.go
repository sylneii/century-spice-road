package main

import "errors"

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
	scoringCards []scoringCard
	goldCoins    int
	silverCoins  int
}

func newScoringStack(scoringCards []scoringCard, goldCoins, silverCoins int) *scoringStack {
	return &scoringStack{
		scoringCards: scoringCards,
		goldCoins:    goldCoins,
		silverCoins:  silverCoins,
	}
}

func (ss *scoringStack) acquire(cardIndex int) error {
	if cardIndex > 4 {
		return errors.New("Card not on display list yet")
	}

}
