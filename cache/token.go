package cache

import "time"

type Token struct {
	accessToken  string
	refreshToken string
	accessExpires      time.Time
	created      time.Time
}

func NewToken(accessToken string, refreshToken string, accessExpires time.Time) *Token {
	return &Token{
		accessToken: accessToken,
		refreshToken: refreshToken,
		accessExpires: accessExpires,
	}
}

func (t *Token) IsExpired() bool {
	return t.accessExpires.Before(time.Now())
}

func (t *Token) GetRefreshToken() string {
	return t.refreshToken
}

func (t *Token) GetAccessToken() string  {
	return t.accessToken
}
