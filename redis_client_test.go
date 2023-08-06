package research_online_redis_go

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisPing(t *testing.T) {
	testPing(t, "redis1:6379")
}

func TestKeydbPing(t *testing.T) {
	testPing(t, "keydb1:6379")
}

func TestDragonflydbPing(t *testing.T) {
	testPing(t, "dragonflydb1:6379")
}

func testPing(t *testing.T, addr string) {
	t.Helper()
	if testing.Short() {
		t.Skip()
	}

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	result, err := client.Ping(context.Background()).Result()
	require.NoError(t, err)
	require.Equal(t, "PONG", result)
}
