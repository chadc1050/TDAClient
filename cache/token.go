package cache

import "time"

type Token struct {
	value   string
	expires time.Time
	created time.Time
}

func NewToken(value string, expiresIn int64) *Token {
	return &Token{
		value:   value,
		created: time.Now(),
		expires: time.Now().Add(time.Duration(expiresIn)),
	}
}

func (t *Token) IsExpired() bool {
	return t.expires.Before(time.Now())
}

func (t *Token) GetValue() string {
	return t.value
}
