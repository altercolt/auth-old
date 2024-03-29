package service

import (
	"context"
	"github.com/altercolt/auth/internal/core/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return UserService{
		repo: repo,
	}
}

func (u UserService) Fetch(ctx context.Context, filter user.Filter) ([]user.User, error) {
	users, err := u.repo.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserService) Create(ctx context.Context, nu user.New) error {
	if errMap := nu.Validate(); errMap != nil {
		return NewValidationError("userService.Create() validation error", errMap)
	}

	passHashByte, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	nu.Password = string(passHashByte)

	model := user.Model{
		Email:     &nu.Email,
		Username:  &nu.Username,
		Firstname: &nu.Firstname,
		Lastname:  &nu.Lastname,
		BirthDate: &nu.BirthDate,
		PassHash:  &nu.Password,
	}

	return u.repo.Create(ctx, &model)
}

func (u UserService) FetchOneByUsername(ctx context.Context, username string) (user.User, error) {
	filter := user.Filter{
		Username: []string{username},
	}

	return u.repo.FetchOne(ctx, filter)
}

func (u UserService) FetchOneByEmail(ctx context.Context, email string) (user.User, error) {
	filter := user.Filter{
		Username: []string{email},
	}

	return u.repo.FetchOne(ctx, filter)
}

func (u UserService) FetchOneByID(ctx context.Context, id int) (user.User, error) {
	filter := user.Filter{
		ID: []int{id},
	}

	return u.repo.FetchOne(ctx, filter)
}

func (u UserService) Update(ctx context.Context, update user.Update, id int) error {
	if errMap := update.Validate(); errMap != nil {
		return NewValidationError("userService.Update() validation error", errMap)
	}

	model := user.Model{
		ID:        &id,
		Email:     update.Email,
		Username:  update.Username,
		Firstname: update.Firstname,
		Lastname:  update.Lastname,
		BirthDate: update.BirthDate,
	}

	if update.Password != nil {
		passHashBytes, err := bcrypt.GenerateFromPassword([]byte(*update.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		passHash := string(passHashBytes)

		model.PassHash = &passHash
	}

	return u.repo.Update(ctx, &model)
}

func (u UserService) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
