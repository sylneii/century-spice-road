package main

import (
	"errors"
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
	points   int
	spices   map[spice]int
}

func (g *Game) AddSpice(s spice, number int) {
	g.currentPlayer.spices[s] += number
}

func (g *Game) RemoveSpice(s spice, number int) error {

	if g.currentPlayer.spices[s] < number {
		slog.Error("err.remove_spice.invalid_amount", "reason", "not enough spice", "spice", s)
		return fmt.Errorf("invalid move: not enough %s", s)
	}
	g.currentPlayer.spices[s] -= number
	return nil
}

func (g *Game) DiscardExtraSpices(discardSpices []spice) {
	for _, spice := range discardSpices {
		g.RemoveSpice(spice, 1)
	}
}
