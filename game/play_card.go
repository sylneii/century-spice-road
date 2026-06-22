package handlers

import (
	"fmt"
	"log/slog"

	"github.com/sylneii/century-spice-road/card"
)

type ActionType string

const  (
	playTradingCard ActionType = "play_card"
	acquireTradingCard ActionType = "acquire_trading_card"
	acquireScoringCard ActionType = "acquire_scoring_card"
	restAction = "rest_action"
)



func (g *GameState) addSpice(s card.Spice, number int64) {
	g.currentPlayer.caravan.spices[s] += number
}

func (g *GameState) removeSpice(s card.Spice, number int64) error {

	if g.currentPlayer.caravan.spices[s] < number {
		slog.Error("err.remove_spice.invalid_amount", "reason", "not enough spice", "spice", s)
		return fmt.Errorf("invalid move: not enough %q", s) //formatverb check
	}
	g.currentPlayer.caravan.spices[s] -= number
	return nil
}

func (g *GameState) discardExtraSpices(discardSpices []card.Spice) {
	for _, spice := range discardSpices {
		g.removeSpice(spice, 1)
	}
}

type actionDetails struct {
	action ActionType	
	cardID string
	position int64
	rest bool	
}

func (g *GameState) ApplyAction(actionType string, actionDetails struct{}) error {
	switch actionType {
	case string(playTradingCard):
		return nil
	case string(acquireTradingCard):
		return nil
	case string(acquireScoringCard):
		return nil
	case string(restAction):
		return nil
	default:
		return nil
	}
}
