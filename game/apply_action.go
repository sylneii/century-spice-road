package game

import (
	"fmt"

	defaultcheck "github.com/gogo/protobuf/test/defaultconflict"
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
	ActionType    actionType `json:"action_type"`
	cardID        string
	position      int64
	spicesPlaced  []card.Spice
	discardSpices []card.Spice
}

func (g *GameState) ApplyAction(actionType actionType, actionDetails actionDetails) error {

	if g.pendingAction != nil {
		switch g.pendingAction.(type) {
		case pendingActionDiscard:
			if actionType != discardSpices {
				return fmt.Errorf("wrong action")
			}
			return g.discardExtraSpices(actionDetails.discardSpices)
		default:

		}
	}

	switch actionDetails.ActionType {
	case playTradingCard:
		return g.playCard(actionDetails.cardID)
	// case string(acquireTradingCard):
	// return g.acquireTradingCard(actionDetails.position)
	// case string(acquireScoringCard):
	// 	return GameState{}, nil
	// case string(restAction):
	// 	return GameState{}, nil
	default:
		return nil
	}
}
