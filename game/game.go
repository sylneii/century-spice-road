package game

import "github.com/sylneii/century-spice-road/cards"

func (g *GameState) getCurrentPlayer() *player {
	return g.players[g.currentPlayerIndex]
}

func (g *GameState) getCurrentPlayerSpices() map[cards.Spice]int64 {
	return g.getCurrentPlayer().caravan.spices
}

func (g *GameState) endTurn() {
	if g.currentPlayerIndex == int64(len(g.players))-1 {
		g.currentPlayerIndex = 0 //turn back to first player
		g.round++
	} else {
		g.currentPlayerIndex++
	}
}
