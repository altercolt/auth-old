package middleware

import (
	"context"
	"errors"
	jwt2 "github.com/altercolt/auth/internal/core/jwt"
	"github.com/altercolt/auth/internal/core/user"
	"github.com/altercolt/auth/pkg/web"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var (
	ErrAuthInvalidToken = errors.New("invalid token")
)

func Auth(a *user.Auth) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			bearer := r.Header.Get("Authorization")
			split := strings.Split(bearer, " ")
			if len(split) != 2 {
				return ErrAuthInvalidToken
			}

			var payload jwt2.Payload

			tkn, err := jwt.ParseWithClaims(split[1], payload, func(token *jwt.Token) (interface{}, error) {

				return nil, nil
			})

			if err != nil {
				return err
			}

			ctx = context.WithValue(ctx, "payload", payload)

			return nil
		}

		return h
	}

	return m
}
