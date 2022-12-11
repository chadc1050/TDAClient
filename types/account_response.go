package types

import "github.com/chadc1050/TDAClient/types/common"

type AccountResponse struct {
	SecuritiesAccount common.SecuritiesAccount `json:"securities_account"`
}
