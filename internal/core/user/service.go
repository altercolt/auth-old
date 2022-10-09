package user

import (
	"context"
)

type Service interface {
	Fetch(ctx context.Context, filter Filter) ([]*User, error)
	Create(ctx context.Context, nu New) error
	FetchOneByUsername(ctx context.Context, username string) (*User, error)
	FetchOneByEmail(ctx context.Context, email string) (*User, error)
	FetchOneByID(ctx context.Context, id int) (*User, error)
	Delete(ctx context.Context, id int) error
}

type Auth interface {
	Login(ctx context.Context, login Login) (*User, error)
	// Authorize Or Authenticate I don't really care tbh
	Authorize(ctx context.Context, token string) ([]string, error)
}
