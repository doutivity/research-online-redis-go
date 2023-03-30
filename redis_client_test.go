package research_online_redis_go

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisPing(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis1:6379",
	})

	result, err := client.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.Equal(t, "PONG", result)
}

func TestKeydbPing(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "keydb1:6379",
	})

	result, err := client.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.Equal(t, "PONG", result)
}

func TestDragonflydbPing(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "dragonflydb1:6379",
	})

	result, err := client.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.Equal(t, "PONG", result)
}
