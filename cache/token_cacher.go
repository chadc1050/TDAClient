package cache

import "sync"

type TokenCacher struct {
	tokenLock sync.Mutex
	token     *Token
}

var (
	instance *TokenCacher
)

func NewTokenCacher(token Token) *TokenCacher {
	if instance == nil {
		instance = &TokenCacher{
			tokenLock: sync.Mutex{},
			token:     &token,
		}
	}

	return instance
}

func (t *TokenCacher) Update(token *Token) {
	t.tokenLock.Lock()

	t.token = token

	t.tokenLock.Unlock()
}

func (t *TokenCacher) IsExpired() bool {
	return t.token.IsExpired()
}

func (t *TokenCacher) GetToken() Token {
	return *t.token
}
