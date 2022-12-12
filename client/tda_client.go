package client

import (
	"encoding/json"
	"errors"
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
	redirectUri string
}

const (
	Environment = "https://api.tdameritrade.com/v1"
)

func NewClient(consumerKey string, redirectUri string) *TDAClient {
	return &TDAClient{
		consumerKey: consumerKey,
		redirectUri: redirectUri,
	}
}

func (client *TDAClient) CheckToken() error {

	if client.t == nil {
		return errors.New("token has not been initialized yet with authentication")
	}

	if client.t.IsExpired() {
		token := client.t.GetToken()
		if err := client.refresh(&token); err != nil {
			return err
		}
	}

	return nil
}

// Authenticate Uses consumer key and code received from OAuth to create a new refresh token and access token.
func (client *TDAClient) Authenticate(oAuthCode string) error {

	authRequest := url.Values{}
	authRequest.Set("grant_type", "authorization_code")
	authRequest.Set("access_type", "offline")
	authRequest.Set("code", oAuthCode)
	authRequest.Set("client_id", client.consumerKey)
	authRequest.Set("redirect_uri", client.redirectUri)

	r, err := http.NewRequest(http.MethodGet, Environment+"/oauth2/token", strings.NewReader(authRequest.Encode()))

	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)

	currentTime := time.Now()

	authResponse := types.AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return err
	}

	token := cache.NewToken(authResponse.AccessToken, authResponse.RefreshToken, currentTime.Add(time.Duration(authResponse.ExpiresIn/1000)))

	client.t = cache.NewTokenCacher(*token)

	return nil
}

func (client *TDAClient) refresh(token *cache.Token) error {

	// TODO: We will also need to handle the refresh token expiring but that will take several months to happen.

	authRequest := url.Values{}
	authRequest.Set("grant_type", "refresh_token")
	authRequest.Set("refresh_token", token.GetRefreshToken())
	authRequest.Set("client_id", client.consumerKey)
	authRequest.Set("redirect_uri", client.redirectUri)

	r, err := http.NewRequest(http.MethodGet, Environment+"/oauth2/token", strings.NewReader(authRequest.Encode()))

	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)

	currentTime := time.Now()

	authResponse := types.AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return err
	}

	updatedToken := cache.NewToken(authResponse.AccessToken, token.GetRefreshToken(), currentTime.Add(time.Duration(authResponse.ExpiresIn/1000)))

	client.t.Update(updatedToken)

	return nil
}
