package auth

import (
	"context"
	"github.com/altercolt/auth/internal/core/user"
)

// Service
// is used for authenticating/authorizing users
type Service interface {
	Login(ctx context.Context, login Login) error
	GenerateTokens(ctx context.Context, usr user.User) (Token, error)
	Authenticate(ctx context.Context, token Token) error
	GetAll(ctx context.Context, userID int) ([]Token, error)
	Revoke(ctx context.Context, userID int, tokenID string) error
	RevokeAll(ctx context.Context, userID int) error
}
