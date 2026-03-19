// Package survey provides the interface to interact with the Redis database. 
package survey

import (
	"fmt"
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

// SaveResponse pushes a form value into a Redis list
func (s *Store) SaveResponse(ctx context.Context, key string, value string) error {
	return s.Client.RPush(ctx, key, value).Err()
}

// GetResults retrieves all entries for a specific question
func (s *Store) GetResults(ctx context.Context, key string) ([]string, error) {
	return s.Client.LRange(ctx, key, 0, -1).Result()
}
