package card

func lookupTradingCard(cardId string) TradingCard {
	return tradingCards[cardId]
}

func GetTradingCard(cardId string) TradingCard {
	return lookupTradingCard(cardId)
}

var tradingCards = map[string]TradingCard{
	"TC01": {
		id:       "TC01",
		CardType: SpiceCard,
		Spices:   map[Spice]int64{turmeric: 2},
	},
	"TC02": {
		id:               "TC02",
		CardType:         UpgradeCard,
		MaxUpgradeLevels: 2,
	},
	"TC03": {
		id:         "TC03",
		CardType:   ExchangeCard,
		FromSpices: map[Spice]int64{turmeric: 2},
		ToSpices:   map[Spice]int64{saffron: 1},
	},
}
