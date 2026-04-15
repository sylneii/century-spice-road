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
	maxUpgradeLevels int
}

func (uc *upgradeCard) play(g *Game, baseSpices []spice, levels []int) {

	var totalLevels int
	for _, level := range levels {
		totalLevels += level
	}
	if totalLevels > uc.maxUpgradeLevels {
		slog.Error("more than max upgrade levels")
		return
	}

	for i, spice := range baseSpices {
		uc.upgradeSpice(g, spice, levels[i])
	}
}

func (uc *upgradeCard) upgradeSpice(g *Game, s spice, levels int) {

	if (s + spice(levels)) > cinnamon {
		slog.Error("spice value greater than cinnamon")
		return
	}
	g.AddSpice(s+spice(levels), 1)
	g.RemoveSpice(s, 1)
}

type exchangeCard struct {
	baseSpice        spice
	exchangeSpice    spice
	ratioNumerator   int64
	ratioDenominator int64
}
