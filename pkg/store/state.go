package store

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

const stateKey = "state"

func (s *Store) GetState(r *http.Request) (string, error) {
	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	state, ok := session.Values[stateKey].(string)
	if !ok {
		return "", errors.New("state token is malformed")
	}

	return state, nil
}

func (s *Store) SaveState(w http.ResponseWriter, r *http.Request) (string, error) {
	state := uuid.New()

	stateData := map[string]interface{}{
		stateKey: state,
	}

	err := s.save(w, r, stateData)
	if err != nil {
		return "", err
	}

	return state.String(), nil
}
