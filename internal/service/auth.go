package service

import (
	"context"
	"github.com/altercolt/auth/internal/core/auth"
	"github.com/altercolt/auth/internal/core/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"time"
)

var (
	// accessTokenDuration
	// accessToken lives for 15 minutes
	accessTokenDuration = time.Minute * 15

	// refreshTokenDuration
	// refreshToken lives for 15 minutes
	refreshTokenDuration = (time.Hour * 24) * 14
)

type AuthService struct {
	userService user.Service
	tokenRepo   auth.TokenRepository
}

func NewAuthService(userService user.Service, tokenRepo auth.TokenRepository) auth.Service {
	return AuthService{
		userService: userService,
		tokenRepo:   tokenRepo,
	}
}

func (a AuthService) Login(ctx context.Context, login auth.Login) error {
	var usr user.User

	_, err := mail.ParseAddress(login.Login)
	if err != nil {
		usr, err = a.userService.FetchOneByUsername(ctx, login.Login)
	} else {
		usr, err = a.userService.FetchOneByEmail(ctx, login.Login)
	}
	if err != nil {
		return nil
	}

	salt := usr.Salt
	passHash := usr.PassHash

	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(login.Password+salt))
	if err != nil {
		return NewAuthorizationError("wrong username/email or password")
	}

	return nil
}

func (a AuthService) GenerateTokens(ctx context.Context, usr user.User) (auth.Token, error) {
	var token auth.Token
	accessTokenPayload := auth.Payload{
		ID:   usr.ID,
		Role: usr.Role,
		Type: auth.AccessToken,
		Exp:  time.Now().Add(15 * time.Minute).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenPayload)
	accessToken, err := access.SignedString("dsa;lkjfa;lksf;alk")
	if err != nil {
		return token, err
	}

	exp := time.Now().Add(15 * time.Minute).Unix()
	refreshTokenPayload := auth.Payload{
		ID:   usr.ID,
		Role: usr.Role,
		Type: auth.RefreshToken,
		Exp:  exp,
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenPayload)
	refreshToken, err := refresh.SignedString("dsalkfjasdjkfa;")
	if err != nil {
		return token, err
	}

	token = auth.Token{
		ID:           uuid.New(),
		UserID:       usr.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExp:    time.Now().Add(accessTokenDuration).Unix(),
		RefreshExp:   time.Now().Add(refreshTokenDuration).Unix(),
	}

	if err = a.tokenRepo.Create(ctx, token); err != nil {
		return auth.Token{}, err
	}

	return token, nil
}

// Authenticate
// Used for validating tokens in middleware.Auth()
// TODO: Finish Authenticate()
func (a AuthService) Authenticate(ctx context.Context, token auth.Token) error {
	access := token.AccessToken
	refresh := token.RefreshToken

	var payload auth.Payload
	tkn, err := jwt.ParseWithClaims(access, &payload,
		func(token *jwt.Token) (interface{}, error) {

			return nil, nil
		})
	if err != nil {
		return err
	}

	payload = tkn.Claims.(auth.Payload)

	return nil
}

func (a AuthService) GetAll(ctx context.Context, userID int) ([]auth.Token, error) {
	return a.tokenRepo.Fetch(ctx, auth.Filter{Users: []int{userID}})
}

func (a AuthService) Revoke(ctx context.Context, userID int, tokenID string) error {
	id, err := uuid.Parse(tokenID)
	if err != nil {
		return NewValidationError("authService.Revoke() uuid validation error", map[string]string{
			"id": "invalid uuid",
		})
	}

	return a.tokenRepo.Delete(ctx, id, userID)
}

func (a AuthService) RevokeAll(ctx context.Context, userID int) error {
	return a.tokenRepo.DeleteAll(ctx, userID)
}
