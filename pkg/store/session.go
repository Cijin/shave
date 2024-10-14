package store

import (
	"errors"
	"log/slog"
	"net/http"

	"shave/pkg/data"

	"github.com/google/uuid"
)

const (
	emailKey       = "email"
	rememberMeKey  = "remember_me"
	accessTokenKey = "access_token"
	userIdKey      = "user_id"
)

// this error does not warrant throwing an internal error, simply
// don't renew the token once expired and log the error
var ErrRememberMe = errors.New("remember me flag is malformed")

func (s *Store) GetSession(r *http.Request) (data.Session, error) {
	sessionData := data.Session{}

	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return sessionData, err
	}

	accessToken, ok := session.Values[accessTokenKey].(string)
	if !ok {
		return sessionData, errors.New("access token is malformed")
	}

	sessionData.AccessToken = accessToken

	return sessionData, nil
}

func (s *Store) GetSessionUser(r *http.Request) (data.SessionUser, error) {
	sessionUser := data.SessionUser{}

	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return sessionUser, err
	}

	email, ok := session.Values[emailKey].(string)
	if !ok {
		return sessionUser, errors.New("email is malformed")
	}

	userId, ok := session.Values[userIdKey].(uuid.UUID)
	if !ok {
		return sessionUser, errors.New("UserId is malformed")
	}

	rememberMe, ok := session.Values[rememberMeKey].(bool)
	if !ok {
		slog.Info("Remember me field is malformed, defaulting to false")

		sessionUser.RememberMe = false
	}

	sessionUser.Email = email
	sessionUser.RememberMe = rememberMe
	sessionUser.UserId = userId

	return sessionUser, nil
}

func (s *Store) SaveSession(w http.ResponseWriter, r *http.Request, sessionData data.Session) error {
	storeData := map[string]interface{}{
		accessTokenKey: sessionData.AccessToken,
	}

	return s.save(w, r, storeData)
}

func (s *Store) SaveSessionUser(w http.ResponseWriter, r *http.Request, sessionUser data.SessionUser) error {
	storeData := map[string]interface{}{
		emailKey:      sessionUser.Email,
		rememberMeKey: sessionUser.RememberMe,
		userIdKey:     sessionUser.UserId,
	}

	return s.save(w, r, storeData)
}
