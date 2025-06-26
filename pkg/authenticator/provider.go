package authenticator

import (
	"context"

	"shave/pkg/data"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Provider interface {
	GetAuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, error)
	ExchangeCode(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	VerifyIdToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error)
	GetUserInfo(idToken *oidc.IDToken) (data.SessionUser, error)
	RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error)
}
