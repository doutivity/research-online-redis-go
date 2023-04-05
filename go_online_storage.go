package research_online_redis_go

import (
	"context"
	"sync"
)

type GoOnlineStorage struct {
	mu   sync.Mutex
	data map[int64]int64
}

func NewGoOnlineStorage() *GoOnlineStorage {
	return &GoOnlineStorage{
		data: map[int64]int64{},
	}
}

func (s *GoOnlineStorage) Store(ctx context.Context, pair UserOnlinePair) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[pair.UserID] = pair.Timestamp

	return nil
}

func (s *GoOnlineStorage) BatchStore(ctx context.Context, pairs []UserOnlinePair) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, pair := range pairs {
		s.data[pair.UserID] = pair.Timestamp
	}

	return nil
}

func (s *GoOnlineStorage) Count(ctx context.Context) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return int64(len(s.data)), nil
}

func (s *GoOnlineStorage) GetAndClear(ctx context.Context) ([]UserOnlinePair, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]UserOnlinePair, 0, len(s.data))
	for userID, timestamp := range s.data {
		result = append(result, UserOnlinePair{
			UserID:    userID,
			Timestamp: timestamp,
		})
	}

	s.data = map[int64]int64{}

	return result, nil
}
