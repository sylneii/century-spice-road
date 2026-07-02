package game

import ()

func (g *GameState) getCurrentPlayer() *player {
	return g.players[g.currentPlayerIndex]
}

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
