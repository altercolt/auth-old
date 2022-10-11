package user

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, nu New) error
	Update(ctx context.Context, update Update, id int) error
	Delete(ctx context.Context, id int) error

	Fetch(ctx context.Context, filter Filter) ([]User, error)
	FetchOneByUsername(ctx context.Context, username string) (User, error)
	FetchOneByEmail(ctx context.Context, email string) (User, error)
	FetchOneByID(ctx context.Context, id int) (User, error)
}
