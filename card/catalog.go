package card

type Spice int64

const (
	turmeric Spice = iota + 1
	saffron
	cardamom
	cinnamon
)

type cardType string

const (
	SpiceCard    cardType = "spice_card"
	UpgradeCard  cardType = "upgrade_card"
	ExchangeCard cardType = "exchange_card"
)

type TradingCard struct {
	id           string
	CardType     cardType
	GetSpices    map[Spice]int64
	UpgradeCount int64
	FromSpices   map[Spice]int64
	ToSpices     map[Spice]int64
}

type ScoringCard struct {
	id        string
	spiceCost map[Spice]int64
	points    int64
}
