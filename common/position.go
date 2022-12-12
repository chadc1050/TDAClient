package common

// Position TODO: Change Float to a BigDecimal
type Position struct {
	ShortQuantity                  float32    `json:"shortQuantity"`
	AveragePrice                   float32    `json:"averagePrice"`
	CurrentDayProfitLoss           float32    `json:"currentDayProfitLoss"`
	CurrentDayProfitLossPercentage float32    `json:"currentDayProfitLossPercentage"`
	LongQuantity                   float32    `json:"long_quantity"`
	SettledLongQuantity            float32    `json:"settledLongQuantity"`
	SettledShortQuantity           float32    `json:"settledShortQuantity"`
	AgedQuantity float32    `json:"agedQuantity"`
	Instrument   Instrument `json:"instrument"`
	MarketValue  float32    `json:"marketValue"`
	MaintenanceRequirement         float32    `json:"maintenanceRequirement"`
	PreviousSessionLongQuantity    float32    `json:"previousSessionLongQuantity"`
}
