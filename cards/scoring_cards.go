package cards

func lookupScoringCard(cardId string) ScoringCard {
	return scoringCards[cardId]
}

func GetScoringCard(cardId string) ScoringCard {
	return lookupScoringCard(cardId)
}

var scoringCards = map[string]ScoringCard{
	"SC01": {
		id:        "SC01",
		spiceCost: map[Spice]int64{turmeric: 3, cardamom: 2},
		points:    6,
	},
}
