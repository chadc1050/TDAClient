package client

import (
	"encoding/json"
	"github.com/chadc1050/TDAClient/types"
	"net/http"
)

type Accounts struct {
	client TDAClient
}

func NewAccounts(t TDAClient) *Accounts {
	return &Accounts{
		client: t,
	}
}

// GetAccount Fields is a comma seperated list of 'orders' and 'positions'
func (accounts *Accounts) GetAccount(accountId string, fields string) (*types.AccountResponse, error) {

	if err := accounts.client.CheckToken(); err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodGet, Environment+"/accounts/"+accountId, nil)

	if err != nil {
		return nil, err
	}

	token := accounts.client.t.GetToken()

	r.Header.Add("Authorization", "Bearer "+token.GetAccessToken())

	if len(fields) != 0 {
		q := r.URL.Query()
		q.Set("fields", fields)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)

	if err != nil {
		return nil, err
	}

	accountsResponse := types.AccountResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&accountsResponse); err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}

func (accounts *Accounts) GetAccounts() {

}
