package game

import (
	"fmt"

	"github.com/sylneii/century-spice-road/cards"
)

func (g *GameState) playCard(details actionDetails) error {
	cardType := cards.GetTradingCardType(details.cardID)

	switch cardType {
	case cards.SpiceCard:
		return g.playSpiceCard(details.cardID)
	case cards.UpgradeCard:
		return g.playUpgradeCard(details.cardID, details.upgradeAction)
		// case card.ExchangeCard:
		// 	return playExchangeCard(g)
	default:
		return nil
	}
}

func (g *GameState) playSpiceCard(cardId string) error {

	//card existence
	card, ok := g.getCurrentPlayer().tradingCards[cardId]
	if !ok {
		return fmt.Errorf("card not in hand")
	}

	for spice, number := range card.Spices {
		g.addSpice(spice, number)
	}

	//check if discard is needed and set pending action to discard to validate against next request from client
	//maintain card in current card till discard is completed and turn is over
	if g.getCurrentPlayer().caravan.isLimitExceeded() {
		g.getCurrentPlayer().currentCard = card
		g.pendingAction = pendingActionDiscard{}
	} else {
		g.restCard(cardId)
		g.endTurn()
	}

	return nil
}

func (g *GameState) playUpgradeCard(cardId string, upgradeDetails upgradeAction) error {

	//card existence
	card, ok := g.getCurrentPlayer().tradingCards[cardId]
	if !ok {
		return fmt.Errorf("card not in hand")
	}

	//validate max upgrade level
	if upgradeDetails.levels > card.MaxUpgradeLevels {
		return fmt.Errorf("not enough upgrades")
	}

	//validate spice upgrade possibility
	if upgradeDetails.spice+cards.Spice(upgradeDetails.levels) > cards.Cinnamon {
		return fmt.Errorf("upgrade value greater than cinnamon not possible")
	}

	//check for spice existence
	if g.getCurrentPlayer().caravan.spices[upgradeDetails.spice] < 1 {
		return fmt.Errorf("requested spice not in caravan")
	}

	//upgrade spice
	g.getCurrentPlayerSpices()[upgradeDetails.spice] -= 1
	g.getCurrentPlayerSpices()[upgradeDetails.spice+cards.Spice(upgradeDetails.levels)] += 1
	if g.pendingActionUpgrade
	//check if upgrades are left
	if 
	return nil
}
