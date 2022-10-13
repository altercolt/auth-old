package auth

import (
	"context"
	"github.com/altercolt/auth/internal/core/user"
)

// Service
// is used for authenticating/authorizing users
type Service interface {
	// Login
	// is used for logging in :)
	Login(ctx context.Context, login Login) error

	// Authenticate
	// is called in middleware.Auth()
	// returns Payload, which is stored in context
	Authenticate(ctx context.Context, accessToken string) (Payload, error)
}

type TokenService interface {
	// GenerateTokens
	// is used for generating new token pair
	GenerateTokens(ctx context.Context, usr user.User) (Token, error)

	// GetAll
	// get all user tokens
	GetAll(ctx context.Context, userID int) ([]Token, error)

	// Revoke
	// is used for revoking single token
	Revoke(ctx context.Context, userID int, tokenID string) error

	// RevokeAll
	// is used for revoking all user tokens
	RevokeAll(ctx context.Context, userID int) error
}
