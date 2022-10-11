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

	salt, passHash, err := u.generatePassword(ctx, nu.Password)
	if err != nil {
		return err
	}

	model := user.Model{
		Email:     &nu.Email,
		Username:  &nu.Username,
		Firstname: &nu.Firstname,
		Lastname:  &nu.Lastname,
		BirthDate: &nu.BirthDate,
		Salt:      &salt,
		PassHash:  &passHash,
	}

	return u.repo.Create(ctx, &model)
}

func (u UserService) generatePassword(ctx context.Context, password string) (string, string, error) {
	salt := u.generateSalt(ctx)
	passHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return salt, string(passHash), nil
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
		salt, passHash, err := u.generatePassword(ctx, *update.Password)
		if err != nil {
			return err
		}

		model.Salt = &salt
		model.PassHash = &passHash
	}

	return u.repo.Update(ctx, &model)
}

func (u UserService) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}

// generateSalt
// TODO: write implementation
func (u UserService) generateSalt(ctx context.Context) string {
	return "salt"
}
