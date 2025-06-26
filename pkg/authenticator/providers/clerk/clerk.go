package clerk

import (
	"context"
	"errors"
	"fmt"
	"os"

	"shave/pkg/data"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Provider struct {
	oidcProvider *oidc.Provider
	config       *oauth2.Config
	issuerURL    string
	callbackURL  string
	clientID     string
}

func New() (*Provider, error) {
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	clerkIssuerURL := os.Getenv("CLERK_ISSUER_URL")
	clerkClientID := os.Getenv("CLERK_CLIENT_ID")
	clerkClientSecret := os.Getenv("CLERK_CLIENT_SECRET")
	host := os.Getenv("HOST")

	if clerkSecretKey == "" || clerkIssuerURL == "" || host == "" || clerkClientID == "" || clerkClientSecret == "" {
		return nil, errors.New("missing required env values for clerk provider")
	}

	callbackURL := fmt.Sprintf("%s/auth/clerk/callback", host)

	clerk.SetKey(clerkSecretKey)
	ctx := context.Background()
	oidcProvider, err := oidc.NewProvider(ctx, clerkIssuerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create oidc provider: %w", err)
	}

	config := &oauth2.Config{
		ClientID:     clerkClientID,
		ClientSecret: clerkClientSecret,
		RedirectURL:  callbackURL,
		Endpoint:     oidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Provider{
		oidcProvider: oidcProvider,
		config:       config,
		issuerURL:    clerkIssuerURL,
		callbackURL:  callbackURL,
		clientID:     clerkClientID,
	}, nil
}

func (p *Provider) GetAuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, error) {
	return p.config.AuthCodeURL(state, opts...), nil
}

func (p *Provider) ExchangeCode(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	token, err := p.config.Exchange(ctx, code, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	return token, nil
}

func (p *Provider) VerifyIdToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: p.clientID,
	}

	return p.oidcProvider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (p *Provider) GetUserInfo(idToken *oidc.IDToken) (data.SessionUser, error) {
	var user data.SessionUser
	var profile map[string]any

	if err := idToken.Claims(&profile); err != nil {
		return user, fmt.Errorf("failed to extract claims: %w", err)
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
		return user, errors.New("idtoken has invalid email verification value")
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

	newToken, err := ts.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return newToken, nil
}

func (p *Provider) VerifyToken(ctx context.Context, sessionToken string) (*clerk.SessionClaims, error) {
	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: sessionToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify session token: %w", err)
	}

	return claims, nil
}
