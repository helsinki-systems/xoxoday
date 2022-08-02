package xoxoday

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

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

type jsonBoolInt bool

func (jbi jsonBoolInt) MarshalJSON() ([]byte, error) {
	var i int

	if jbi {
		i = 1
	} else {
		i = 0
	}

	return json.Marshal(i)
}

func (jbi *jsonBoolInt) UnmarshalJSON(b []byte) error {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	if i != 0 {
		*jbi = jsonBoolInt(true)
	} else {
		*jbi = jsonBoolInt(false)
	}

	return nil
}

type jsonUnixTime time.Time

func (jut jsonUnixTime) MarshalJSON() ([]byte, error) {
	s := strconv.FormatInt(time.Time(jut).UnixMilli(), 10)

	return json.Marshal(s)
}

func (jut *jsonUnixTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	ms, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	t := time.UnixMilli(ms)
	*jut = jsonUnixTime(t)

	return nil
}
