package user

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, m *Model) error
	Update(ctx context.Context, m *Model) error
	Fetch(ctx context.Context, f Filter) ([]*Model, error)
	FetchOne(ctx context.Context, f Filter) (*Model, error)
	Delete(ctx context.Context, id string) error
}

type Filter struct {
	ID        []string
	Email     []string
	Username  []string
	Firstname []string
	Lastname  []string
	BirthDate []time.Time
	Password  []string
}
