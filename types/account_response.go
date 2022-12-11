package types

import "TDAClient/types/common"

type AccountResponse struct {
	SecuritiesAccount common.SecuritiesAccount `json:"securities_account"`
}
