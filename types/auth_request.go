package types

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	AccessType   string `json:"access_type"`
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
}
