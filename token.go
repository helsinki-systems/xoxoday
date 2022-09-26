package xoxoday

type Token struct {
	Success            jsonBoolInt  `json:"success"`
	Status             int          `json:"status"`
	AccessToken        string       `json:"access_token"`
	TokenType          string       `json:"token_type"`
	ExpiresIn          int          `json:"expires_in,string"`
	RefreshToken       string       `json:"refresh_token"`
	AccessTokenExpiry  jsonUnixTime `json:"access_token_expiry"`
	RefreshTokenExpiry jsonUnixTime `json:"refresh_token_expiry"`
}
