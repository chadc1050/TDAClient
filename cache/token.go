package cache

import "time"

type Token struct {
	value   string
	expires time.Time
	created time.Time
}

func NewToken(value string, expires time.Time) *Token {
	return &Token{
		value:   value,
		expires: expires,
	}
}

func (t *Token) IsExpired() bool {
	return t.expires.Before(time.Now())
}

func (t *Token) GetValue() string {
	return t.value
}
