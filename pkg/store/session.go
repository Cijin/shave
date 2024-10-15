package store

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net/http"

	"shave/pkg/data"

	"github.com/google/uuid"
)

const (
	emailKey         = "email"
	subKey           = "sub"
	emailVerifiedKey = "email_verified"
	nameKey          = "name"
	avatarURLKey     = "avatar_url"
	userIdKey        = "user_id"
	accessTokenKey   = "access_token"
	expiresAtKey     = "expires_at"
)

func (s *Store) GetSession(r *http.Request) (data.Session, error) {
	sessionData := data.Session{}

	session, err := s.cookieStore.Get(r, sessionName)
	if err != nil {
		return sessionData, err
	}

	compressedAccessToken, ok := session.Values[accessTokenKey].([]byte)
	if !ok {
		return sessionData, errors.New("access token is malformed")
	}

	rData := bytes.NewReader(compressedAccessToken)
	zr, err := gzip.NewReader(rData)
	if err != nil {
		return sessionData, err
	}

	accessToken, err := io.ReadAll(zr)
	if err != nil {
		return sessionData, err
	}

	sessionData.AccessToken = string(accessToken)

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

	name, ok := session.Values[nameKey].(string)
	if !ok {
		return sessionUser, errors.New("name is malformed")
	}

	avatarURL, ok := session.Values[avatarURLKey].(string)
	if !ok {
		return sessionUser, errors.New("avatar URL is malformed")
	}

	emailVerified, ok := session.Values[emailVerifiedKey].(bool)
	if !ok {
		return sessionUser, errors.New("email verified is malformed")
	}

	sub, ok := session.Values[subKey].(string)
	if !ok {
		return sessionUser, errors.New("sub is malformed")
	}

	userId, ok := session.Values[userIdKey].(uuid.UUID)
	if !ok {
		return sessionUser, errors.New("UserId is malformed")
	}

	sessionUser.Email = email
	sessionUser.AvatarURL = avatarURL
	sessionUser.Name = name
	sessionUser.Sub = sub
	sessionUser.EmailVerified = emailVerified
	sessionUser.UserId = userId

	return sessionUser, nil
}

func (s *Store) SaveSession(w http.ResponseWriter, r *http.Request, sessionData data.Session) error {
	var b bytes.Buffer
	zw := gzip.NewWriter(&b)
	zw.Name = "access-token"
	if _, err := zw.Write([]byte(sessionData.AccessToken)); err != nil {
		return err
	}

	if err := zw.Flush(); err != nil {
		return err
	}

	if err := zw.Close(); err != nil {
		return err
	}

	storeData := map[string]interface{}{
		accessTokenKey: b,
		expiresAtKey:   sessionData.ExpiresAt,
	}

	return s.save(w, r, storeData)
}

func (s *Store) SaveSessionUser(w http.ResponseWriter, r *http.Request, sessionUser data.SessionUser) error {
	storeData := map[string]interface{}{
		emailKey:         sessionUser.Email,
		subKey:           sessionUser.Sub,
		userIdKey:        sessionUser.UserId,
		emailVerifiedKey: sessionUser.EmailVerified,
		avatarURLKey:     sessionUser.AvatarURL,
		nameKey:          sessionUser.Name,
	}

	return s.save(w, r, storeData)
}
