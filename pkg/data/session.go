package data

import (
	"context"
	"net/mail"

	"github.com/google/uuid"
)

type Session struct {
	AccessToken string
}

func (s Session) Valid(ctx context.Context) Problems {
	problems := NewProblems()

	if s.AccessToken == "" {
		problems.Add("AccessToken", "Acess token is empty")
	}

	return problems
}

type SessionUser struct {
	UserId     uuid.UUID
	Email      string
	RememberMe bool
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
	State    string
}

func (sv SessionVerifier) Valid(ctx context.Context) Problems {
	problems := NewProblems()

	if sv.State == "" {
		problems.Add("state", "State is empty")
	}

	if sv.State == "" {
		problems.Add("state", "State is empty")
	}

	return problems
}
