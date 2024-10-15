package store

import (
	"errors"
	"net/http"

	"shave/pkg/data"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

const (
	stateKey    = "state"
	verifierKey = "verifier"
)

func (s *Store) GetSessionVerfier(r *http.Request) (data.SessionVerifier, error) {
	var sessionData data.SessionVerifier

	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return sessionData, err
	}

	state, ok := session.Values[stateKey].(uuid.UUID)
	if !ok {
		return sessionData, errors.New("state is malformed")
	}

	verifier, ok := session.Values[verifierKey].(string)
	if !ok {
		return sessionData, errors.New("verifier is malformed")
	}

	return data.SessionVerifier{State: state, Verifier: verifier}, nil
}

func (s *Store) SaveSessionVerfier(w http.ResponseWriter, r *http.Request) (data.SessionVerifier, error) {
	state := uuid.New()
	verifier := oauth2.GenerateVerifier()

	stateData := map[string]interface{}{
		stateKey:    state,
		verifierKey: verifier,
	}

	err := s.save(w, r, stateData)
	if err != nil {
		return data.SessionVerifier{}, err
	}

	return data.SessionVerifier{State: state, Verifier: verifier}, nil
}
