package types

type AuthRequest struct {
	grantType    string 'json:"grant_type"'
	refreshToken string
	accessType   string
	code         string
	clientId     string
	redirectUri  string
}
