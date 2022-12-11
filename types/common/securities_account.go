package common

type SecuritiesAccount struct {
	Type                    string          `json:"type"`
	AccountID               string          `json:"accountId"`
	RoundTrips              int32           `json:"roundTrips"`
	IsDayTrader             bool            `json:"isDayTrader"`
	IsClosingOnlyRestricted bool            `json:"isClosingOnlyRestricted"`
	Positions               []Position      `json:"positions"`
	OrderStrategies         []OrderStrategy `json:"orderStrategies"`
}
