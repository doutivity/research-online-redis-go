package research_online_redis_go

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

// 10 minutes
const defaultOnlineSetGroup = 600

type SetOnlineStorage struct {
	client *redis.Client
	group  int64
}

func NewSetOnlineStorage(client *redis.Client, group int64) *SetOnlineStorage {
	if group == 0 {
		group = defaultOnlineSetGroup
	}

	return &SetOnlineStorage{client: client, group: group}
}

func (s *SetOnlineStorage) Store(ctx context.Context, pair UserOnlinePair) error {
	return s.client.SAdd(ctx, s.key(pair.Timestamp), pair.UserID).Err()
}

func (s *SetOnlineStorage) Count(ctx context.Context) (int64, error) {
	keys, err := s.keys(ctx)
	if err != nil {
		return 0, err
	}

	count := int64(0)
	for _, key := range keys {
		groupCount, err := s.client.SCard(ctx, key).Result()
		if err != nil {
			return 0, err
		}

		count += groupCount
	}

	return count, nil
}

func (s *SetOnlineStorage) GetAndClear(ctx context.Context) ([]UserOnlinePair, error) {
	keys, err := s.keys(ctx)
	if err != nil {
		return nil, err
	}

	var result []UserOnlinePair
	for _, key := range keys {
		oldKey := key
		newKey := "s:online:tmp"

		err := s.client.Rename(ctx, oldKey, newKey).Err()
		if err != nil {
			return result, err
		}

		timestamp, err := s.parseKey(oldKey)
		if err != nil {
			return result, err
		}

		userIDs, err := s.client.SMembers(ctx, newKey).Result()
		if err != nil {
			return result, err
		}

		for _, stringUserID := range userIDs {
			userID, err := strconv.ParseInt(stringUserID, 10, 64)
			if err != nil {
				// unreachable, ignore for article

				// logging or use https://github.com/hashicorp/go-multierror

				// just in case
				continue
			}

			result = append(result, UserOnlinePair{
				UserID:    userID,
				Timestamp: timestamp,
			})
		}
	}

	return result, nil
}

func (s *SetOnlineStorage) keys(ctx context.Context) ([]string, error) {
	return s.client.Keys(ctx, "s:online:main:*").Result()
}

func (s *SetOnlineStorage) key(timestamp int64) string {
	return fmt.Sprintf("s:online:main:%d", (timestamp/s.group)*s.group)
}

func (s *SetOnlineStorage) parseKey(key string) (timestamp int64, err error) {
	_, err = fmt.Fscanf(strings.NewReader(key), "s:online:main:%d", &timestamp)
	if err != nil {
		return 0, err
	}

	return timestamp, nil
}
