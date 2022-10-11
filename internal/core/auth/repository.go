package auth

import (
	"context"
	"github.com/google/uuid"
)

// TokenRepository
// is used for storing tokens in the database
type TokenRepository interface {
	Create(ctx context.Context, token Token) error
	Fetch(ctx context.Context, filter Filter) ([]Token, error)
	Delete(ctx context.Context, id uuid.UUID, userID int) error
	DeleteAll(ctx context.Context, userID int) error
}

type Filter struct {
	IDs           []uuid.UUID
	Users         []int
	AccessTokens  []string
	RefreshTokens []string
}
