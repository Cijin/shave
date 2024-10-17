package google

import (
	"context"
	"errors"
	"fmt"
	"os"

	"shave/pkg/data"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	issuerURL    = "https://accounts.google.com"
	providerName = "google"
)

type Provider struct {
	config       oauth2.Config
	oidcProvider *oidc.Provider
	name         string
}

func New() (*Provider, error) {
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
		RedirectURL:  fmt.Sprintf("%s/auth/%s/callback", callbackDomain, providerName),
		Endpoint:     google.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Provider{name: providerName, config: conf, oidcProvider: oidcProvider}, nil
}

func (p *Provider) GetName() string {
	return p.name
}

func (p *Provider) GetAuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return p.config.AuthCodeURL(state, opts...)
}

func (p *Provider) ExchangeCode(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return p.config.Exchange(ctx, code, opts...)
}

func (p *Provider) VerifyIdToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: p.config.ClientID,
	}

	return p.oidcProvider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (p *Provider) GetUserInfo(idToken *oidc.IDToken) (data.SessionUser, error) {
	var user data.SessionUser
	var profile map[string]interface{}

	if err := idToken.Claims(&profile); err != nil {
		return user, err
	}

	name, ok := profile["name"].(string)
	if !ok {
		return user, errors.New("idtoken has invalid name type")
	}

	avatarURL, ok := profile["picture"].(string)
	if !ok {
		return user, errors.New("idtoken has avatar url of invalid type")
	}

	email, ok := profile["email"].(string)
	if !ok {
		return user, errors.New("idtoken has invalid email")
	}

	emailVerified, ok := profile["email_verified"].(bool)
	if !ok {
		return user, errors.New("idtoken has invalid email verifiation value")
	}

	sub, ok := profile["sub"].(string)
	if !ok {
		return user, errors.New("idtoken has invalid sub value")
	}

	user.Name = name
	user.AvatarURL = avatarURL
	user.Email = email
	user.EmailVerified = emailVerified
	user.Sub = sub

	return user, nil
}

func (p *Provider) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{RefreshToken: refreshToken}
	ts := p.config.TokenSource(ctx, token)

	return ts.Token()
}
