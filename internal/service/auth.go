package service

import (
	"context"
	"github.com/altercolt/auth/internal/core/auth"
	"github.com/altercolt/auth/internal/core/user"
	"github.com/altercolt/auth/pkg/keystore"
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

// AuthService
// implements auth.Service
type AuthService struct {
	userService user.Service
	keyStore    keystore.KeyStore
}

func NewAuthService(userService user.Service) auth.Service {
	return AuthService{
		userService: userService,
	}
}

func (a AuthService) Authenticate(ctx context.Context, accessToken string) (auth.Payload, error) {
	//TODO implement me
	panic("implement me")
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

	passHash := usr.PassHash

	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(login.Password))
	if err != nil {
		return NewAuthorizationError("wrong username/email or password")
	}

	return nil
}

// TokenService
// implements auth.TokenService
type TokenService struct {
	tokenRepo auth.TokenRepository
}

func NewTokenService(tokenRepo auth.TokenRepository) auth.TokenService {
	return TokenService{
		tokenRepo: tokenRepo,
	}
}

func (t TokenService) GenerateTokens(ctx context.Context, usr user.User) (auth.Token, error) {
	var token auth.Token
	accessTokenPayload := auth.Payload{
		ID:   usr.ID,
		Role: usr.Role,
		Type: auth.AccessToken,
		Exp:  time.Now().Add(15 * time.Minute).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenPayload)
	//TODO
	accessToken, err := access.SignedString("dsa;lkjfa;lksf;alk")
	if err != nil {
		return auth.Token{}, err
	}

	exp := time.Now().Add(15 * time.Minute).Unix()
	refreshTokenPayload := auth.Payload{
		ID:   usr.ID,
		Role: usr.Role,
		Type: auth.RefreshToken,
		Exp:  exp,
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenPayload)
	//TODO
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

	if err = t.tokenRepo.Create(ctx, token); err != nil {
		return auth.Token{}, err
	}

	return token, nil
}

func (t TokenService) RefreshTokens(ctx context.Context, token auth.Token) (auth.Payload, error) {

	return auth.Payload{}, nil
}

func (t TokenService) GetAll(ctx context.Context, userID int) ([]auth.Token, error) {
	return t.tokenRepo.Fetch(ctx, auth.Filter{Users: []int{userID}})
}

func (t TokenService) Revoke(ctx context.Context, userID int, tokenID string) error {
	id, err := uuid.Parse(tokenID)
	if err != nil {
		return NewValidationError("authService.Revoke() uuid validation error", map[string]string{
			"id": "invalid uuid",
		})
	}
	return t.tokenRepo.Delete(ctx, id, userID)
}

func (t TokenService) RevokeAll(ctx context.Context, userID int) error {
	return t.tokenRepo.DeleteAll(ctx, userID)
}
