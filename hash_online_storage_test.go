package research_online_redis_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedisHashOnlineStorage(t *testing.T) {
	testHashOnlineStorage(t, "redis1:6379")
}

func TestKeydbHashOnlineStorage(t *testing.T) {
	testHashOnlineStorage(t, "keydb1:6379")
}

func TestDragonflydbHashOnlineStorage(t *testing.T) {
	testHashOnlineStorage(t, "dragonflydb1:6379")
}

func testHashOnlineStorage(t *testing.T, addr string) {
	t.Helper()

	ctx := context.Background()

	client, err := Client(ctx, addr)
	require.NoError(t, err)

	require.NoError(t, client.FlushAll(ctx).Err())

	storage := NewHashOnlineStorage(client)

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
