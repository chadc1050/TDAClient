package client

import (
	"encoding/json"
	"github.com/chadc1050/TDAClient/cache"
	"github.com/chadc1050/TDAClient/types"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type TDAClient struct {
	t           *cache.TokenCacher
	consumerKey string
	oAuthCode   string
}

const (
	Environment = "https://api.tdameritrade.com/v1/"
)

func NewClient(consumerKey string, oAuthCode string) *TDAClient {

	var token, err = authorize(consumerKey, oAuthCode)

	if err != nil {
		panic("An error occurred while authorizing!")
	}

	return &TDAClient{
		consumerKey: consumerKey,
		oAuthCode:   oAuthCode,
		t:           cache.NewTokenCacher(*token),
	}

}

func (client *TDAClient) CheckToken() error {
	if client.t.IsExpired() {
		token := client.t.GetToken()
		newToken, err := authorize(client.consumerKey, token.GetValue())
		if err != nil {
			return err
		}

		client.t.Update(newToken)
	}

	return nil
}

func authorize(consumerKey string, oAuthCode string) (*cache.Token, error) {

	authRequest := url.Values{}
	authRequest.Set("grant_type", "refresh_token")
	authRequest.Set("code", oAuthCode)
	authRequest.Set("client_id", consumerKey)

	r, err := http.NewRequest(http.MethodGet, Environment+"/oauth2/token", strings.NewReader(authRequest.Encode()))

	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)

	currentTime := time.Now()

	authResponse := types.AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	return cache.NewToken(
		authResponse.RefreshToken,
		currentTime.Add(time.Duration(authResponse.RefreshTokenExpiresIn/1000))), nil
}
