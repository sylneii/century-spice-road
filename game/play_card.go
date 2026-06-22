package game

import "github.com/sylneii/century-spice-road/card"

func (g *GameState) playCard(cardId string) error {
	playedCard := card.GetTradingCard(cardId)

	switch playedCard.CardType {
	case card.SpiceCard:
		return playSpiceCard(g, playedCard)
		// case card.UpgradeCard:
		// 	return playUpgradeCard(g)
		// case card.ExchangeCard:
		// 	return playExchangeCard(g)
	}

	return nil
}

func playSpiceCard(g *GameState, card card.TradingCard) error {
	for spice, number := range card.GetSpices {
		g.addSpice(spice, number)
	}
	return nil
}
