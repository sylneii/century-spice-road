package cards

func lookupTradingCardType(cardId string) cardType {
	return tradingCards[cardId].CardType
}

func GetTradingCardType(cardId string) cardType {
	return lookupTradingCardType(cardId)
}

var tradingCards = map[string]TradingCard{
	"TC01": {
		Id:       "TC01",
		CardType: SpiceCard,
		Spices:   map[Spice]int64{turmeric: 2},
	},
	"TC02": {
		Id:               "TC02",
		CardType:         UpgradeCard,
		MaxUpgradeLevels: 2,
	},
	"TC03": {
		Id:         "TC03",
		CardType:   ExchangeCard,
		FromSpices: map[Spice]int64{turmeric: 2},
		ToSpices:   map[Spice]int64{saffron: 1},
	},
}
