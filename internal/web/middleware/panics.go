package middleware

import (
	"context"
	"github.com/altercolt/auth/pkg/web"
	"net/http"
)

func Panics() web.Middleware {

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			return nil
		}

		return h
	}

	return m
}
