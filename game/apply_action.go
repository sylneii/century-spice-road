package game

import (
	"github.com/sylneii/century-spice-road/card"
)

type actionType string

const (
	playTradingCard    actionType = "play_card"
	acquireTradingCard actionType = "acquire_trading_card"
	acquireScoringCard actionType = "acquire_scoring_card"
	restAction         actionType = "rest_action"

	discardSpices actionType = "discard_spices"
	endTurn       actionType = "end_turn"
)

type actionDetails struct {
	actionType    string
	cardID        string
	position      int64
	spicesPlaced  []card.Spice
	discardSpices []card.Spice
}

func (g *GameState) ApplyAction(actionType string, actionDetails actionDetails) error {
	switch actionDetails.actionType {
	case string(playTradingCard):
		return g.playCard(actionDetails.cardID)
	// case string(acquireTradingCard):
	// return g.acquireTradingCard(actionDetails.position)
	// case string(acquireScoringCard):
	// 	return GameState{}, nil
	// case string(restAction):
	// 	return GameState{}, nil
	case string(discardSpices):
		return g.discardExtraSpices(actionDetails.discardSpices)
	default:
		return nil
	}
}
