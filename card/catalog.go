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
	spiceCard cardType = "spice_card"
	upgradeCard cardType = "upgrade_card"
	exchangeCard cardType = "exchange_card"
)

type TradingCard struct {
	id string
	cardType cardType
	getSpices map[Spice]int64
	upgradeCount int64
	fromSpices map[Spice]int64
	toSpices map[Spice]int64
}

type ScoringCard struct {
	id string
	spiceCost map[Spice]int64
	points int64
}




