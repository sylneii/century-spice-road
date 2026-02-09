package main

import (
	"log/slog"

	"golang.org/x/text/unicode/rangetable"
)

type spiceCard struct {
	spices map[spice]int
}

func (sc *spiceCard) play(g *Game) {
	for spice, number := range sc.spices {
		g.AddSpice(spice, number)
	}
}

type upgradeCard struct {
	upgradeLevels int
}

func (uc *upgradeCard) play(g *Game) {
	remainingUpgradeLevels := uc.upgradeLevels
	for range remainingUpgradeLevels {

	}
}

func (uc *upgradeCard) upgradeSpice(g *Game, s spice) {

	g.RemoveSpice(s, 1)

	switch s {
	case turmeric:
		g.AddSpice(saffron, 1)
	case saffron:
		g.AddSpice(cardamom, 1)
	case cardamom:
		g.AddSpice(cinnamon, 1)
	}
}
