package mongodb

import "github.com/altercolt/auth/internal/core/user"

type UserRepository struct {
}

func NewUserRepository() user.Repository {
	return &UserRepository{}
}
