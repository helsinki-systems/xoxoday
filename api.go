package xoxoday

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type API struct {
	env env

	client *http.Client
}

func (a *API) makeURL(p string) string {
	return a.baseURL() + p
}

func (a *API) baseURL() string {
	switch a.env {
	case EnvDevelopment:
		return "https://stagingaccount.xoxoday.com/chef"
	case EnvProduction:
		return "https://accounts.xoxoday.com/chef"
	}

	panic("unknown env, this must never occur")
}

func (a *API) Run(r Request) ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	res, err := a.client.Post(
		a.makeURL("/v1/oauth/api"),
		"application/json",
		bytes.NewReader(j),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		var e Error
		if err := json.Unmarshal(body, &e); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}

		return nil, e
	}

	return body, nil
}

type OAuthConfig oauth2.Config

func New(env env, t Token, oac OAuthConfig) *API {
	c := oauth2.Config(oac)

	return &API{
		env: env,

		client: (&c).Client(context.Background(), &oauth2.Token{
			AccessToken:  t.AccessToken,
			TokenType:    t.TokenType,
			RefreshToken: t.RefreshToken,
			Expiry:       time.Time(t.AccessTokenExpiry),
		}),
	}
}
