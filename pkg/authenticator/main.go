package authenticator

import (
	"context"
	"errors"
	"net/http"

	"shave/pkg/data"

	"golang.org/x/oauth2"
)

type Authenticator struct {
	provider           Provider
	shouldRefreshToken bool
}

func New(shouldRefreshToken bool, provider Provider) *Authenticator {
	return &Authenticator{provider, shouldRefreshToken}
}

func (a *Authenticator) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, error) {
	if a.shouldRefreshToken {
		opts = append(opts, oauth2.AccessTypeOffline)
	}

	return a.provider.GetAuthCodeURL(state, opts...)
}

func (a *Authenticator) Authenticate(r *http.Request, opts ...oauth2.AuthCodeOption) (*oauth2.Token, data.SessionUser, error) {
	var user data.SessionUser

	code := r.URL.Query().Get("code")
	if code == "" {
		return nil, user, errors.New("code is empty")
	}

	token, err := a.provider.ExchangeCode(r.Context(), code, opts...)
	if err != nil {
		return nil, user, err
	}

	if !token.Valid() {
		return nil, user, errors.New("token recieved is invalid")
	}

	idToken, err := a.provider.VerifyIdToken(r.Context(), token)
	if err != nil {
		return nil, user, err
	}

	user, err = a.provider.GetUserInfo(idToken)
	if err != nil {
		return nil, user, err
	}

	return token, user, nil
}

func (a *Authenticator) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	return a.provider.RefreshToken(ctx, refreshToken)
}
