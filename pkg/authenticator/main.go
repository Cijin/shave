package authenticator

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	providers map[string]Provider
}

func New(providers ...Provider) *Authenticator {
	p := make(map[string]Provider, len(providers))

	for _, provider := range providers {
		p[provider.GetName()] = provider
	}

	return &Authenticator{p}
}

func (a *Authenticator) getProvider(r *http.Request) (Provider, error) {
	providerName := chi.URLParam(r, "provider")
	provider, ok := a.providers[providerName]
	if !ok {
		return nil, fmt.Errorf("Provider:'%s' is not a registered provider", providerName)
	}

	return provider, nil
}

func (a *Authenticator) AuthCodeURL(state string, r *http.Request) (string, error) {
	provider, err := a.getProvider(r)
	if err != nil {
		return "", err
	}

	return provider.GetAuthCodeURL(state), nil
}

func (a *Authenticator) Authenticate(r *http.Request) (*oauth2.Token, *oidc.IDToken, error) {
	provider, err := a.getProvider(r)
	if err != nil {
		return nil, nil, err
	}

	code := chi.URLParam(r, "code")
	if code == "" {
		return nil, nil, errors.New("code is empty")
	}

	token, err := provider.ExchangeCode(r.Context(), code)
	if err != nil {
		return nil, nil, err
	}

	if !token.Valid() {
		return nil, nil, errors.New("token recieved is invalid")
	}

	idToken, err := provider.VerifyIssuer(r.Context(), token)
	if err != nil {
		return nil, nil, err
	}

	return token, idToken, nil
}
