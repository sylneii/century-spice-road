package game

import ()

func (g *GameState) getCurrentPlayer() *player {
	return g.players[g.currentPlayerIndex]
}

func (g *GameState) endTurn() {
	if g.currentPlayerIndex == int64(len(g.players))-1 {
		g.currentPlayerIndex = 0 //turn back to first player
		g.round++
	} else {
		g.currentPlayerIndex++
	}
}
