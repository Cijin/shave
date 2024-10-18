package handlers

import (
	"context"
	"testing"
	"time"

	"shave/internal/database"
	"shave/pkg/data"
)

func TestGetOrCreateUser(t *testing.T) {
	t.Setenv("ENV", "TEST")
	t.Setenv("SESSION_SECRET", "TEST_SESSION_SECRET")

	h, err := NewHttpHandler(db)
	if err != nil {
		t.Error("unable to create handler", err)
	}

	tests := map[string]struct {
		userExists   bool
		sessionUser  data.SessionUser
		existingUser database.CreateUserParams
		expectedUser database.User
	}{
		"Should return existing user": {
			userExists: true,
			existingUser: database.CreateUserParams{
				ID:            getUUID().String(),
				Email:         "test@email.com",
				Sub:           "test-sub",
				Name:          "test-user",
				EmailVerified: true,
				CreatedAt:     time.Now().UTC(),
				UpdatedAt:     time.Now().UTC(),
			},
			sessionUser: data.SessionUser{Email: "test@email.com", Sub: "test-sub", Name: "test-user", EmailVerified: true},
		},
		"Should create a new user": {
			userExists: false,
			existingUser: database.CreateUserParams{
				ID:            getUUID().String(),
				Email:         "test@email.com",
				Sub:           "test-sub",
				Name:          "test-user",
				EmailVerified: true,
				CreatedAt:     time.Now().UTC(),
				UpdatedAt:     time.Now().UTC(),
			},
			sessionUser: data.SessionUser{Email: "test1@email.com", Sub: "test-sub-2", Name: "test-user", EmailVerified: true},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testUser, err := h.dbQueries.CreateUser(context.Background(), test.existingUser)
			if err != nil {
				t.Error("unable to add user to db", err)
			}

			user, err := h.getOrCreateUser(context.Background(), test.sessionUser)
			if err != nil {
				t.Errorf("expected no error but got=%v", err)
			}

			if test.userExists {
				if testUser.ID != user.ID {
					t.Errorf("expected user id=%s, got=%s", testUser.ID, user.ID)
				}
			} else {
				if testUser.ID == user.ID {
					t.Errorf("extected new user to be created but expected id=%s, got=%s", testUser.ID, user.ID)
				}
			}

			err = h.dbQueries.DeleteUser(context.Background(), testUser.ID)
			if err != nil {
				t.Error("unable to delete test user:", err)
			}
		})
	}
}
