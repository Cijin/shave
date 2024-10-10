package store

import (
	"errors"
	"log/slog"
	"net/http"

	"shave/pkg/data"

	"github.com/google/uuid"
)

const (
	EmailKey       = "email"
	RememberMeKey  = "remember_me"
	AccessTokenKey = "access_token"
	UserIdKey      = "user_id"
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

	accessToken, ok := session.Values[AccessTokenKey].(string)
	if !ok {
		return sessionData, errors.New("access token is malformed")
	}

	sub, ok := session.Values[SubKey].(string)
	if !ok {
		return sessionData, errors.New("sub is malformed")
	}

	sessionData.AccessToken = accessToken
	sessionData.Sub = sub

	return sessionData, nil
}

func (s *Store) GetSessionUser(r *http.Request) (data.SessionUser, error) {
	sessionUser := data.SessionUser{}

	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return sessionUser, err
	}

	email, ok := session.Values[EmailKey].(string)
	if !ok {
		return sessionUser, errors.New("email is malformed")
	}

	userId, ok := session.Values[UserIdKey].(uuid.UUID)
	if !ok {
		return sessionUser, errors.New("UserId is malformed")
	}

	rememberMe, ok := session.Values[RememberMeKey].(bool)
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
		AccessTokenKey: sessionData.AccessToken,
		SubKey:         sessionData.Sub,
	}

	return s.save(w, r, storeData)
}

func (s *Store) SaveSessionUser(w http.ResponseWriter, r *http.Request, sessionUser data.SessionUser) error {
	storeData := map[string]interface{}{
		EmailKey:      sessionUser.Email,
		RememberMeKey: sessionUser.RememberMe,
		RoleKey:       sessionUser.Role,
		UserIdKey:     sessionUser.UserId,
	}

	return s.save(w, r, storeData)
}
