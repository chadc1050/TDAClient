package client

import (
	"TDAClient/cache"
	"encoding/json"
	"io"
	"net/http"
)

type TDAClient struct {
	t *cache.TokenCacher

	id  string
	key string
}

const Environment = "https://api.tdameritrade.com/v1/"

func NewClient(id string, key string) *TDAClient {

	var token = authorize()

	return &TDAClient{
		id:  id,
		key: key,
		t:   cache.NewTokenCacher(token),
	}
}

func (c *TDAClient) reAuthorize(t Token) {
	http.

}

func authorize() (*cache.Token, error) {

	r, err := http.NewRequest("", Environment+"/oauth2/token")
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(io.ReadAll(resp.Body), &)
}
