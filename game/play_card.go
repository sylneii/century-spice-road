package game

import (
	"fmt"

	"github.com/sylneii/century-spice-road/card"
)

func (g *GameState) playCard(cardId string) error {
	cardType := card.GetTradingCardType(cardId)

	switch cardType {
	case card.SpiceCard:
		return g.playSpiceCard(cardId)
		// case card.UpgradeCard:
		// 	return playUpgradeCard(g)
		// case card.ExchangeCard:
		// 	return playExchangeCard(g)
	default:
		return nil
	}
}

func (g *GameState) playSpiceCard(cardId string) error {

	card, ok := g.getCurrentPlayer().tradingCards[cardId]
	if !ok {
		return fmt.Errorf("card not in hand")
	}

	for spice, number := range card.Spices {
		g.addSpice(spice, number)
	}

	if g.getCurrentPlayer().caravan.isLimitExceeded() {
		g.getCurrentPlayer().currentCard = card
		g.pendingAction = pendingActionDiscard{}
	} else {
		g.restCard(cardId)
		g.endTurn()
	}

	return nil
}

func playUpgradeCard(g *GameState, card card.TradingCard) error {

	return nil
}
