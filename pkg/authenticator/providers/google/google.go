package google

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const issuerURL = "https://accounts.google.com/"

type Provider struct {
	config       oauth2.Config
	oidcProvider *oidc.Provider
	name         string
}

func New(name string) (*Provider, error) {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackDomain := os.Getenv("CALLBACK_DOMAIN")

	if clientId == "" || clientSecret == "" || callbackDomain == "" {
		return nil, errors.New("missing required env values for google provider")
	}

	oidcProvider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  fmt.Sprintf("%s/auth/%s/callback", callbackDomain, name),
		Endpoint:     google.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Provider{name: "google", config: conf, oidcProvider: oidcProvider}, nil
}

func (p *Provider) GetName() string {
	return p.name
}

func (p *Provider) GetAuthCodeURL(state string) string {
	return p.config.AuthCodeURL(state)
}

func (p *Provider) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return p.config.Exchange(ctx, code)
}

func (p *Provider) VerifyIdIssuer(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: p.config.ClientID,
	}

	return p.oidcProvider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
