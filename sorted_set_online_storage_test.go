package research_online_redis_go

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

var sortedSetOnlineStorageConstructor = func(client *redis.Client) OnlineStorage {
	return NewSortedSetOnlineStorage(client)
}

func TestRedisSortedSetOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "redis1:6379", sortedSetOnlineStorageConstructor)
}

func TestKeydbSortedSetOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "keydb1:6379", sortedSetOnlineStorageConstructor)
}

func TestDragonflydbSortedSetOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "dragonflydb1:6379", sortedSetOnlineStorageConstructor)
}

func BenchmarkRedisSortedSetOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "redis1:6379", sortedSetOnlineStorageConstructor)
}

func BenchmarkKeydbSortedSetOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "keydb1:6379", sortedSetOnlineStorageConstructor)
}

func BenchmarkDragonflydbSortedSetOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "dragonflydb1:6379", sortedSetOnlineStorageConstructor)
}
