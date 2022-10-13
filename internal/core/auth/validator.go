package auth

import (
	"errors"
	"time"
)

var (
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenInvalidType = errors.New("invalid token type")
)

// Valid
// implements jwt.Claims Interface
func (p Payload) Valid() error {
	if p.Type != RefreshToken && p.Type != AccessToken {
		return ErrTokenInvalidType
	}

	exp := time.Unix(p.Exp, 0)
	if exp.After(time.Now()) {
		return ErrTokenExpired
	}

	return nil
}
