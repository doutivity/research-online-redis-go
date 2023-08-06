package research_online_redis_go

import (
	"context"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGoOnlineStorage(t *testing.T) {
	testMapOnlineStorage(t, NewGoOnlineStorage())
}

func BenchmarkGoOnlineStorage(b *testing.B) {
	benchmarkMapOnlineStorage(b, NewGoOnlineStorage())
}

func testMapOnlineStorage(
	t *testing.T,
	storage OnlineStorage,
) {
	t.Helper()

	ctx := context.Background()

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

func benchmarkMapOnlineStorage(
	b *testing.B,
	storage OnlineStorage,
) {
	b.Helper()

	ctx := context.Background()

	var (
		startTimestamp = time.Now().Unix()
		startUserID    = int64(1e7)
	)

	b.ResetTimer()
	if os.Getenv("MODE") == "parallel" {
		var (
			counter = int64(0)
		)

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
	require.Equal(b, int64(b.N), actualCount)
}
