package research_online_redis_go

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type SortedSetOnlineStorage struct {
	client *redis.Client
}

func NewSortedSetOnlineStorage(client *redis.Client) *SortedSetOnlineStorage {
	return &SortedSetOnlineStorage{client: client}
}

func (s *SortedSetOnlineStorage) Store(ctx context.Context, pair UserOnlinePair) error {
	return s.client.ZAdd(ctx, "z:online:main", redis.Z{
		Score:  float64(pair.Timestamp),
		Member: strconv.FormatInt(pair.UserID, 10),
	}).Err()
}

func (s *SortedSetOnlineStorage) Count(ctx context.Context) (int64, error) {
	return s.client.ZCard(ctx, "z:online:main").Result()
}

func (s *SortedSetOnlineStorage) GetAndClear(ctx context.Context) ([]UserOnlinePair, error) {
	var (
		oldKey = "z:online:main"
		newKey = "z:online:tmp"
	)

	err := s.client.Rename(ctx, oldKey, newKey).Err()
	if err != nil {
		return nil, err
	}

	members, err := s.client.ZRangeWithScores(ctx, newKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	result := make([]UserOnlinePair, 0, len(members))
	for _, member := range members {
		stringUserID, ok := member.Member.(string)
		if !ok {
			// unreachable, ignore for article

			// just in case
			continue
		}

		userID, err := strconv.ParseInt(stringUserID, 10, 64)
		if err != nil {
			// unreachable, ignore for article

			// logging or use https://github.com/hashicorp/go-multierror

			// just in case
			continue
		}

		result = append(result, UserOnlinePair{
			UserID:    userID,
			Timestamp: int64(member.Score),
		})
	}

	return result, nil
}
