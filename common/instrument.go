package common

type Instrument struct {
	AssetType   string `json:"assetType"`
	Cusip       string `json:"cusip"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}
