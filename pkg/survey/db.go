// Package survey provides the interface to interact with the Redis database. 
package survey

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	Client *redis.Client
}

func NewStore(addr string) *Store {
	return &Store{
		Client: redis.NewClient(&redis.Options{Addr: addr}),
	}
}

func (s *Store) SaveResponse(ctx context.Context, key string, value string) error {
	return s.Client.RPush(ctx, key, value).Err()
}

func (s *Store) GetResults(ctx context.Context, key string) ([]string, error) {
	return s.Client.LRange(ctx, key, 0, -1).Result()
}
