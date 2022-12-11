package common

type OrderLeg struct {
	OrderLegType   string     `json:"orderLegType"`
	LegId          int32      `json:"legId"`
	Instrument     Instrument `json:"instrument"`
	Instruction    string     `json:"instruction"`
	PositionEffect string     `json:"positionEffect"`
	Quantity       float32    `json:"quantity"`
	QuantityType   string     `json:"quantityType"`
}
