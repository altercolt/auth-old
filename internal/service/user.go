package service

import "github.com/altercolt/auth/internal/core/user"

type UserService struct {
}

func NewUserService() user.Service {
	return &UserService{}
}
