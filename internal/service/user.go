package service

import (
	"context"
	"github.com/altercolt/auth/internal/core/user"
)

type UserService struct {
}

func NewUserService() user.Service {
	return UserService{}
}

func (u UserService) Fetch(ctx context.Context, filter user.Filter) ([]*user.User, error) {
	return nil, nil
}

func (u UserService) Create(ctx context.Context, nu user.New) error {
	return nil
}

func (u UserService) FetchOneByUsername(ctx context.Context, username string) (*user.User, error) {
	return nil, nil
}

func (u UserService) FetchOneByEmail(ctx context.Context, email string) (*user.User, error) {
	return nil, nil
}

func (u UserService) FetchOneByID(ctx context.Context, id int) (*user.User, error) {
	return nil, nil
}

func (u UserService) Delete(ctx context.Context, id int) error {
	return nil
}
