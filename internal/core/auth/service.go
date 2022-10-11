package auth

import "context"

type Auth interface {
	Login(ctx context.Context, login Login) (Token, error)
	Authenticate(ctx context.Context, token Token) error
}
