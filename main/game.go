package main

import (
	"fmt"
	"log/slog"
)

type Game struct {
	players       []player
	currentPlayer player
	round         int
}

type player struct {
	username string
	caravan  caravan
}

type caravan struct {
	spices       map[spice]int
	scoringCards []scoringCard
}

func (g *Game) AddSpice(s spice, number int) {
	g.currentPlayer.caravan.spices[s] += number
}

func (g *Game) RemoveSpice(s spice, number int) error {

	if g.currentPlayer.caravan.spices[s] < number {
		slog.Error("err.remove_spice.invalid_amount", "reason", "not enough spice", "spice", s)
		return fmt.Errorf("invalid move: not enough %q", s) //formatverb check
	}
	g.currentPlayer.caravan.spices[s] -= number
	return nil
}

func (g *Game) DiscardExtraSpices(discardSpices []spice) {
	for _, spice := range discardSpices {
		g.RemoveSpice(spice, 1)
	}
}
