package user

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, m *Model) error
	Update(ctx context.Context, m *Model) error
	Fetch(ctx context.Context, f Filter) ([]User, error)
	FetchOne(ctx context.Context, f Filter) (User, error)
	Delete(ctx context.Context, id int) error
}

type Filter struct {
	ID        []int
	Email     []string
	Username  []string
	Firstname []string
	Lastname  []string
	BirthDate []time.Time
}
