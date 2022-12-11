package types

type AuthResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int32  `json:"expires_in"`
	Scope                 string `json:"scope"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}
