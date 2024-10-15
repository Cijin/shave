package data

import (
	"context"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	AccessToken string
	ExpiresAt   time.Time
}

func (s Session) Valid(ctx context.Context) Problems {
	problems := NewProblems()

	if s.AccessToken == "" {
		problems.Add("AccessToken", "Acess token is empty")
	}

	return problems
}

type SessionUser struct {
	UserId        uuid.UUID
	Sub           string
	AvatarURL     string
	EmailVerified bool
	Name          string
	Email         string
}

func (su SessionUser) Valid(ctx context.Context) Problems {
	problems := NewProblems()

	_, err := mail.ParseAddress(su.Email)
	if err != nil {
		problems.Add("Email", err.Error())
	}

	return problems
}

type SessionVerifier struct {
	Verifier string
	State    uuid.UUID
}
