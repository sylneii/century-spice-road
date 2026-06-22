package card

func lookupTradingCard(cardId string) TradingCard {
	return tradingCards[cardId]
}

func GetTradingCard(cardId string) TradingCard {
	return lookupTradingCard(cardId)
}

var tradingCards = map[string]TradingCard{
	"TC01": {
		id: "TC01",
		cardType: spiceCard,
		getSpices: map[Spice]int64{turmeric: 2},
	},
	"TC02": {
		id: "TC02",
		cardType: upgradeCard,
		upgradeCount: 2,
	},
	"TC03": {
		id: "TC03",
		cardType: exchangeCard,
		fromSpices: map[Spice]int64{turmeric: 2},
		toSpices: map[Spice]int64{saffron: 1},
	},
}

