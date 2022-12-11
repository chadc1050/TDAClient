package types

import "TDAClient/types/common"

type AccountResponse struct {
	securitiesAccount common.SecuritiesAccount `json:"securities_account"`
}
