package postgres

import (
	"context"

	"github.com/altercolt/auth/internal/core/user"
)

type UserRepository struct {
}

func (u UserRepository) Create(ctx context.Context, m *user.Model) error {

	return nil
}

func (u UserRepository) Update(ctx context.Context, m *user.Model) error {

	return nil
}

func (u UserRepository) Fetch(ctx context.Context, f user.Filter) ([]*user.Model, error) {

	return nil, nil
}

func (u UserRepository) FetchOne(ctx context.Context, f user.Filter) (*user.Model, error) {

	return nil, nil
}

func (u UserRepository) Delete(ctx context.Context, id string) error {

	return nil
}

func NewUserRepository() user.Repository {

	return nil
}
