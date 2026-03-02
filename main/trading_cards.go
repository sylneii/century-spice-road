package main

import "log/slog"

type spiceCard struct {
	spices map[spice]int
}

func (sc *spiceCard) play(g *Game) {
	for spice, number := range sc.spices {
		g.AddSpice(spice, number)
	}
}

type upgradeCard struct {
	maxUpgradeLevels       int
	remainingUpgradeLevels int
}

func (uc *upgradeCard) play(g *Game) {

}

func (uc *upgradeCard) upgradeSpice(g *Game, s spice, levels int) {

	if (s + spice(levels)) > cinnamon {
		slog.Error("invalid move")
		return
	}
	g.AddSpice(s+spice(levels), 1)
	g.RemoveSpice(s, 1)
}
