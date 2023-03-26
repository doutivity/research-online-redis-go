package research_online_redis_go

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

type onlineStorageConstructor func(client *redis.Client) OnlineStorage

func testOnlineStorage(t *testing.T, addr string, newStorage onlineStorageConstructor) {
	t.Helper()

	ctx := context.Background()

	client, err := Client(ctx, addr)
	require.NoError(t, err)

	require.NoError(t, client.FlushDB(ctx).Err())

	storage := newStorage(client)

	expected := []UserOnlinePair{
		{
			UserID:    1000001,
			Timestamp: 1679800725,
		},
		{
			UserID:    1000002,
			Timestamp: 1679800730,
		},
		{
			UserID:    1000003,
			Timestamp: 1679800735,
		},
	}

	for _, pair := range expected {
		err := storage.Store(ctx, pair)

		require.NoError(t, err)
	}

	actual, err := storage.GetAndClear(ctx)
	require.NoError(t, err)

	requireUserOnlinePairsEqual(t, expected, actual)
}
