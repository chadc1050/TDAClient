package common

import "time"

type OrderStrategy struct {
	Session                 string          `json:"session"`
	Duration                string          `json:"duration"`
	OrderType               string          `json:"orderType"`
	CancelTime              Date            `json:"cancelTime"`
	ComplexStrategyType     string          `json:"complexStrategyType"`
	Quantity                float32         `json:"quantity"`
	FilledQuantity          float32         `json:"filledQuantity"`
	RemainingQuantity       float32         `json:"remainingQuantity"`
	RequestedDestination    string          `json:"requestedDestination"`
	DestinationLinkName     string          `json:"destinationLinkName"`
	ReleaseTime             time.Time       `json:"releaseTime"`
	StopPrice               float32         `json:"stopPrice"`
	StopPriceLinkBasis      string          `json:"stopPriceLinkBasis"`
	StopPriceLinkType       string          `json:"stopPriceLinkType"`
	StopPriceOffset         float32         `json:"stopPriceOffset"`
	StopType                string          `json:"stopType"`
	PriceLinkBasis          string          `json:"priceLinkBasis"`
	PriceLinkType           string          `json:"priceLinkType"`
	Price                   float32         `json:"price"`
	TaxLotMethod            string          `json:"taxLotMethod"`
	OrderLegCollection      []OrderLeg      `json:"orderLegCollection"`
	ActivationPrice         float32         `json:"activationPrice"`
	SpecialInstruction      string          `json:"specialInstruction"`
	OrderStrategyType       string          `json:"orderStrategyType"`
	OrderId                 int64           `json:"orderId"`
	Cancellable             bool            `json:"cancellable"`
	Editable                bool            `json:"editable"`
	Status                  string          `json:"status"`
	EnteredTime             time.Time       `json:"enteredTime"`
	CloseTime               time.Time       `json:"closeTime"`
	Tag                     string          `json:"tag"`
	AccountId               int64           `json:"accountId"`
	OrderActivityCollection []OrderActivity `json:"orderActivityCollection"`
	//TODO: There are two more fields here that weren't clear
	StatusDescription string `json:"statusDescription"`
}
