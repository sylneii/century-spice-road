package game

import "github.com/sylneii/century-spice-road/card"

func (g *GameState) playCard(cardId string) error {
	playedCard := card.GetTradingCard(cardId)

	switch playedCard.CardType {
	case card.SpiceCard:
		return g.playSpiceCard(playedCard)
		// case card.UpgradeCard:
		// 	return playUpgradeCard(g)
		// case card.ExchangeCard:
		// 	return playExchangeCard(g)
	default:
		return nil
	}
}

func (g *GameState) playSpiceCard(card card.TradingCard) error {
	for spice, number := range card.Spices {
		g.addSpice(spice, number)
	}
	if g.getCurrentPlayer().caravan.isLimitExceeded() {
		g.pendingAction = &pendingActionDiscard{}
	} else {
		g.phase = phaseEndTurn
	}

	return nil
}

func playUpgradeCard(g *GameState, card card.TradingCard) error {

}
