package authenticator

import (
	"context"

	"shave/pkg/data"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Provider interface {
	GetName() string
	GetAuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	ExchangeCode(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	VerifyIssuer(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error)
	GetUserInfo(idToken *oidc.IDToken) (data.SessionUser, error)
}
