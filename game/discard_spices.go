package game

import (
	"fmt"

	"github.com/sylneii/century-spice-road/cards"
)

func (g *GameState) discardExtraSpices(discardSpices []card.Spice) error {
	discardMap := make(map[card.Spice]int64)
	for _, spice := range discardSpices {
		discardMap[spice]++
		if g.getCurrentPlayer().caravan.spices[spice] < discardMap[spice] {
			return fmt.Errorf("invalid move: not enough %q", spice) //format verb check
		}
	}

	for spice, number := range discardMap {
		g.removeSpice(spice, number)
	}

	g.restCard(g.getCurrentPlayer().currentCard.Id)
	g.getCurrentPlayer().currentCard = card.TradingCard{}

	return nil
}
