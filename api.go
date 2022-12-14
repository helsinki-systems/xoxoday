package xoxoday

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

func New(
	ctx context.Context,
	env env,
	t Token,
	oac OAuthConfig,
) *API {
	verifyToken(t)

	ac := &API{
		env: env,
	}

	oac.Endpoint = oauth2.Endpoint{
		TokenURL: ac.makeURL("/v1/oauth/token/" + t.TokenType),
		AuthURL:  ac.makeURL("/v1/oauth/token"),
	}
	c := oauth2.Config(oac)

	ac.client = (&c).Client(ctx, &oauth2.Token{
		AccessToken:  t.AccessToken,
		TokenType:    t.TokenType,
		RefreshToken: t.RefreshToken,
		Expiry:       t.AccessTokenExpiry.Time,
	})

	return ac
}
