package data

import (
	"context"
	"net/mail"
	"slices"
	"time"

	"github.com/google/uuid"
)

var SupportedProviders = []string{"google", "github"}

type Session struct {
	AccessToken string
	Expiry      time.Time
	Provider    string
}

func (s Session) Valid(ctx context.Context) Problems {
	problems := NewProblems()

	if s.AccessToken == "" {
		problems.Add("AccessToken", "Acess token is empty")
	}

	if !slices.Contains(SupportedProviders, s.Provider) {
		problems.Add("Provider", "Provider is invalid")
	}

	if s.Expiry.IsZero() {
		problems.Add("Expiry", "Expiry is invalid")
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

	if su.UserId == uuid.Nil {
		problems.Add("UserId", "UserId is an empty UUID")
	}

	if su.Sub == "" {
		problems.Add("Sub", "Sub is empty")
	}

	if su.AvatarURL == "" {
		problems.Add("AvatarURL", "AvatarURL is empty")
	}

	if su.Name == "" {
		problems.Add("Name", "Name is empty")
	}

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
