package main

import (
	"fmt"
	"log/slog"
)

type GameState struct {
	players       []player
	currentPlayer player
	round         int
	scoringStack
	tradingStack
}

type player struct {
	username     string
	caravan      caravan
	tradingCards []tradingCard
	scoringCards []scoringCard
	goldCoins    int64
	silverCoins  int64
}

type caravan struct {
	spices map[spice]int64
}

func (g *GameState) AddSpice(s spice, number int64) {
	g.currentPlayer.caravan.spices[s] += number
}

func (g *GameState) RemoveSpice(s spice, number int64) error {

	if g.currentPlayer.caravan.spices[s] < number {
		slog.Error("err.remove_spice.invalid_amount", "reason", "not enough spice", "spice", s)
		return fmt.Errorf("invalid move: not enough %q", s) //formatverb check
	}
	g.currentPlayer.caravan.spices[s] -= number
	return nil
}

func (g *GameState) DiscardExtraSpices(discardSpices []spice) {
	for _, spice := range discardSpices {
		g.RemoveSpice(spice, 1)
	}
}

func (g *GameState) ApplyAction(actionType string, actionDetails struct{}) {

}
