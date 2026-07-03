package game

import (
	"fmt"

	"github.com/sylneii/century-spice-road/card"
)

func (c caravan) isLimitExceeded() bool {
	var total int64
	for _, count := range c.spices {
		total += count
	}

	if total > caravanLimit {
		return true
	}

	return false
}

func (g *GameState) addSpice(s card.Spice, number int64) {
	g.getCurrentPlayer().caravan.spices[s] += number
}

func (g *GameState) removeSpice(s card.Spice, number int64) error {

	if g.getCurrentPlayer().caravan.spices[s] < number {
		return fmt.Errorf("invalid move: not enough %q", s) //formatverb check
	}
	g.getCurrentPlayer().caravan.spices[s] -= number
	return nil
}

// use up the played trading card and add it to the rest pile
func (g *GameState) restCard(cardId string) {
	g.getCurrentPlayer().usedCards[cardId] = g.getCurrentPlayer().tradingCards[cardId]
	delete(g.getCurrentPlayer().tradingCards, cardId)
}
