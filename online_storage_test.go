package research_online_redis_go

import (
	"context"
	"os"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

type onlineStorageConstructor func(client *redis.Client) OnlineStorage

func testOnlineStorage(t *testing.T, addr string, newStorage onlineStorageConstructor) {
	t.Helper()

	ctx := context.Background()

	client, err := Client(ctx, addr)
	require.NoError(t, err)

	require.NoError(t, client.FlushAll(ctx).Err())

	storage := newStorage(client)

	expected := []UserOnlinePair{
		{
			UserID:    10000001,
			Timestamp: 1679800725,
		},
		{
			UserID:    10000002,
			Timestamp: 1679800730,
		},
		{
			UserID:    10000003,
			Timestamp: 1679800735,
		},
	}

	for _, pair := range expected {
		err := storage.Store(ctx, pair)

		require.NoError(t, err)
	}

	actualCount, err := storage.Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(len(expected)), int64(actualCount))

	actual, err := storage.GetAndClear(ctx)
	require.NoError(t, err)

	requireUserOnlinePairsEqual(t, expected, actual)
}

func benchmarkOnlineStorage(b *testing.B, addr string, newStorage onlineStorageConstructor) {
	b.Helper()

	ctx := context.Background()

	client, err := Client(ctx, addr)
	require.NoError(b, err)

	require.NoError(b, client.FlushDB(ctx).Err())

	storage := newStorage(client)

	var (
		expectedCount  = int64(b.N)
		startTimestamp = time.Now().Truncate(time.Hour).Unix()
		startUserID    = int64(1e7)
	)

	b.ResetTimer()

	if os.Getenv("MODE") == "parallel" {
		var (
			counter = int64(0)
		)

		if os.Getenv("BATCH") == "" {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					index := atomic.AddInt64(&counter, 1)

					err := storage.Store(ctx, UserOnlinePair{
						UserID:    startUserID + index,
						Timestamp: startTimestamp + index,
					})

					require.NoError(b, err)
				}
			})
		} else {
			batch, err := strconv.ParseInt(os.Getenv("BATCH"), 10, 64)
			require.NoError(b, err)
			require.True(b, batch >= 1)

			expectedCount *= batch

			b.RunParallel(func(pb *testing.PB) {
				pairs := make([]UserOnlinePair, batch)

				for pb.Next() {
					index := atomic.AddInt64(&counter, batch)

					for i := int64(0); i < batch; i++ {
						pairs[i] = UserOnlinePair{
							UserID:    startUserID + index - i,
							Timestamp: startTimestamp + index - i,
						}
					}

					err := storage.BatchStore(ctx, pairs)

					require.NoError(b, err)
				}
			})
		}
	} else {
		for index := int64(0); index < int64(b.N); index++ {
			err := storage.Store(ctx, UserOnlinePair{
				UserID:    startUserID + index,
				Timestamp: startTimestamp + index,
			})

			require.NoError(b, err)
		}
	}

	actualCount, err := storage.Count(ctx)
	require.NoError(b, err)
	require.Equal(b, expectedCount, actualCount)
}
