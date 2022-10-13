package middleware

import (
	"context"
	"github.com/altercolt/auth/internal/core/auth"
	"github.com/altercolt/auth/internal/service"
	"github.com/altercolt/auth/pkg/web"
	"net/http"
	"strings"
)

func Auth(a auth.Service) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			bearer := r.Header.Get("Authorization")
			split := strings.Split(bearer, " ")
			if len(split) != 2 {
				return service.NewAuthorizationError("invalid token")
			}

			payload, err := a.Authenticate(ctx, split[1])
			if err != nil {
				return service.NewAuthorizationError("invalid token: " + err.Error())
			}

			ctx = context.WithValue(ctx, "payload", payload)

			return nil
		}

		return h
	}

	return m
}
