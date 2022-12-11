package client

import (
	"bytes"
	"encoding/json"
	"github.com/chadc1050/TDAClient/cache"
	"github.com/chadc1050/TDAClient/types"
	"net/http"
	"time"
)

type TDAClient struct {
	t   *cache.TokenCacher
	id  string
	key string
}

const (
	Environment = "https://api.tdameritrade.com/v1/"
)

func NewClient(id string, key string) *TDAClient {

	var token, err = authorize(id, key)

	if err != nil {
		panic("An error occurred while authorizing!")
	}

	return &TDAClient{
		id:  id,
		key: key,
		t:   cache.NewTokenCacher(*token),
	}

}

func (client *TDAClient) CheckToken() error {
	if client.t.IsExpired() {
		token := client.t.GetToken()
		newToken, err := client.reAuthorize(token)
		if err != nil {
			return err
		}

		client.t.Update(newToken)
	}

	return nil
}

func (client *TDAClient) reAuthorize(t cache.Token) (*cache.Token, error) {

	authRequest := types.AuthRequest{
		GrantType:   "refresh_token",
		AccessType:  "offline",
		RedirectUri: "http://localhost:8080/",
	}

	b := bytes.Buffer{}

	if err := json.NewEncoder(&b).Encode(authRequest); err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodGet, Environment+"/oauth2/token", &b)

	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)

	currentTime := time.Now()

	authResponse := types.AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	return cache.NewToken(
		authResponse.AccessToken,
		authResponse.RefreshToken,
		currentTime.Add(time.Duration(authResponse.RefreshTokenExpiresIn/1000))), nil
}

func authorize(id string, key string) (*cache.Token, error) {

	authRequest := types.AuthRequest{
		GrantType:   "authorization_code",
		AccessType:  "offline",
		Code:        key,
		ClientId:    id,
		RedirectUri: "http://localhost:8080/",
	}

	b := bytes.Buffer{}

	if err := json.NewEncoder(&b).Encode(authRequest); err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodGet, Environment+"/oauth2/token", &b)

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
		authResponse.AccessToken,
		authResponse.RefreshToken,
		currentTime.Add(time.Duration(authResponse.RefreshTokenExpiresIn/1000))), nil
}
