package usergrp

import (
	"context"
	"encoding/json"
	"github.com/altercolt/auth/internal/core/user"
	"net/http"
)

type Handler struct {
	service user.Service
}

func (h Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Handler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu user.New
	if err := json.NewDecoder(r.Body).Decode(&nu); err != nil {
		return err
	}

	if err := h.service.Create(ctx, nu); err != nil {
		return err
	}

	return nil
}

func (h Handler) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h Handler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Token
// used for getting new access-refresh token pair
func (h Handler) Token(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Login
// For logging in :)
func (h Handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return nil
}
