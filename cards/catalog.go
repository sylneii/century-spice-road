package cards

type Spice int64

const (
	turmeric Spice = iota + 1
	saffron
	cardamom
	Cinnamon
)

type cardType string

const (
	SpiceCard    cardType = "spice_card"
	UpgradeCard  cardType = "upgrade_card"
	ExchangeCard cardType = "exchange_card"
)

type TradingCard struct {
	Id               string
	CardType         cardType
	Spices           map[Spice]int64
	MaxUpgradeLevels int64
	FromSpices       map[Spice]int64
	ToSpices         map[Spice]int64
}

type ScoringCard struct {
	id        string
	spiceCost map[Spice]int64
	points    int64
}
