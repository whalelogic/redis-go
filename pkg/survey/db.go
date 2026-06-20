// Package survey provides the interface to interact with the Redis database. 
package survey

import (
	"context"
	"sort"
	"github.com/redis/go-redis/v9"
)

// Store provides methods to interact with the Redis database for storing and retrieving survey responses.
type Store struct {
	Client *redis.Client
}

// SurveyResult represents a survey response and its tally. 
type SurveyResult struct {
	Value string
	Count int64
}

// NewStore initializes a new Store.
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

// GetSurveyResults gets the survey results and returns the tally of each response, sorted by count descending.
func (s *Store) GetSurveyResults(ctx context.Context, key string) ([]SurveyResult, error) {
	values, err := s.GetResults(ctx, key)
	if err != nil {
		return nil, err
	}

	resultMap := make(map[string]int64)
	for _, value := range values {
		resultMap[value]++
	}

	results := make([]SurveyResult, 0, len(resultMap))
	for value, count := range resultMap {
		results = append(results, SurveyResult{Value: value, Count: count})
	}
	
	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})
	
	return results, nil
}






