package main

type caravan struct {
	turmeric    int
	saffron     int
	cardamom    int
	clove       int
	totalSpices int
}

func (c *caravan) addSpice(spice int, number int) {
	c.turmeric = number
}
