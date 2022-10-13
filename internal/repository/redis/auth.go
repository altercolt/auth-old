package redis

import (
	"context"
	"github.com/altercolt/auth/internal/core/auth"
	"github.com/altercolt/auth/internal/repository"
	"github.com/go-redis/redis/v9"
	"time"
)

type TokenStore struct {
	db *redis.Client
}

func NewTokenStore(db *redis.Client) auth.TokenStore {
	return TokenStore{
		db: db,
	}
}

func (t TokenStore) Set(ctx context.Context, exp time.Duration, key string, value ...string) error {
	status := t.db.Set(ctx, key, value, exp)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (t TokenStore) Get(ctx context.Context, key string) ([]string, error) {
	var token []string

	res := t.db.Get(ctx, key)

	err := res.Scan(&token)

	if err == redis.Nil {
		return nil, repository.NewNotFoundError("tokenStore.Get() error: not found")
	} else if err != nil && err != redis.Nil {
		return nil, err
	}

	return token, nil
}
