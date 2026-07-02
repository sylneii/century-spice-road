package game

import (
	"fmt"

	"github.com/sylneii/century-spice-road/card"
)

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
