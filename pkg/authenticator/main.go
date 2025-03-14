package authenticator

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"shave/pkg/data"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	providers          map[string]Provider
	shouldRefreshToken bool
}

func New(shouldRefreshToken bool, providers ...Provider) *Authenticator {
	p := make(map[string]Provider, len(providers))

	for _, provider := range providers {
		p[provider.GetName()] = provider
	}

	return &Authenticator{p, shouldRefreshToken}
}

func (a *Authenticator) getProvider(r *http.Request) (Provider, error) {
	providerName := chi.URLParam(r, "provider")
	provider, ok := a.providers[providerName]
	if !ok {
		return nil, fmt.Errorf("Provider:'%s' is not a registered provider", providerName)
	}

	return provider, nil
}

func (a *Authenticator) AuthCodeURL(state string, r *http.Request, opts ...oauth2.AuthCodeOption) (string, error) {
	provider, err := a.getProvider(r)
	if err != nil {
		return "", err
	}

	if a.shouldRefreshToken {
		opts = append(opts, oauth2.AccessTypeOffline)
	}

	return provider.GetAuthCodeURL(state, opts...), nil
}

func (a *Authenticator) Authenticate(r *http.Request, opts ...oauth2.AuthCodeOption) (*oauth2.Token, data.SessionUser, error) {
	var user data.SessionUser
	provider, err := a.getProvider(r)
	if err != nil {
		return nil, user, err
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		return nil, user, errors.New("code is empty")
	}

	token, err := provider.ExchangeCode(r.Context(), code, opts...)
	if err != nil {
		return nil, user, err
	}

	if !token.Valid() {
		return nil, user, errors.New("token recieved is invalid")
	}

	idToken, err := provider.VerifyIdToken(r.Context(), token)
	if err != nil {
		return nil, user, err
	}

	user, err = provider.GetUserInfo(idToken)
	if err != nil {
		return nil, user, err
	}

	return token, user, nil
}

func (a *Authenticator) RefreshToken(ctx context.Context, providerName, refreshToken string) (*oauth2.Token, error) {
	provider, ok := a.providers[providerName]
	if !ok {
		return nil, fmt.Errorf("Provider:'%s' is not a registered provider", providerName)
	}

	return provider.RefreshToken(ctx, refreshToken)
}
