package authgrp

import (
	"context"
	"github.com/altercolt/auth/internal/core/auth"
	"net/http"
)

type Handler struct {
	service auth.Service
}

func (h Handler) GetTokens(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Handler) Revoke(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Handler) RevokeAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Handler) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
