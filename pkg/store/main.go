package store

import (
	"encoding/gob"
	"errors"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

const sessionName = "paper-chase"

// required to store user id of type UUID into
// session cookie
func init() {
	gob.Register(uuid.UUID{})
}

type Store struct {
	cookieStore *sessions.CookieStore
}

func New() (*Store, error) {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		return nil, errors.New("session secret has not been set")
	}

	cookieStore := sessions.NewCookieStore([]byte(secret))
	cookieStore.Options.Path = "/"
	cookieStore.Options.Secure = true
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.SameSite = http.SameSiteStrictMode

	return &Store{cookieStore}, nil
}

func (s *Store) Update(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values[key] = value

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

// use methods to access save, ensures data saved is consistent with app requirements
func (s *Store) save(w http.ResponseWriter, r *http.Request, values map[string]interface{}) error {
	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return err
	}

	for k, v := range values {
		session.Values[k] = v
	}

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Clear(w http.ResponseWriter, r *http.Request) error {
	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values = make(map[interface{}]interface{})

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

// used for testing only
func (s *Store) GetValues(w http.ResponseWriter, r *http.Request) (map[interface{}]interface{}, error) {
	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return nil, err
	}

	return session.Values, nil
}
