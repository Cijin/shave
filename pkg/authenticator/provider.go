package authenticator

import (
	"context"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Provider interface {
	GetName() string
	GetAuthCodeURL(state string) string
	ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error)
	VerifyIssuer(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error)
}
