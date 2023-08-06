package research_online_redis_go

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

var (
	setOnlineStorageTestConstructor = func(client *redis.Client) OnlineStorage {
		return NewSetOnlineStorage(client, 1)
	}
	setOnlineStorageBenchmarkConstructor = func(client *redis.Client) OnlineStorage {
		return NewSetOnlineStorage(client, 1800)
	}
)

func TestRedisSetOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "redis1:6379", setOnlineStorageTestConstructor)
}

func TestKeydbSetOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "keydb1:6379", setOnlineStorageTestConstructor)
}

func TestDragonflydbSetOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "dragonflydb1:6379", setOnlineStorageTestConstructor)
}

func BenchmarkRedisSetOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "redis1:6379", setOnlineStorageBenchmarkConstructor)
}

func BenchmarkKeydbSetOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "keydb1:6379", setOnlineStorageBenchmarkConstructor)
}

func BenchmarkDragonflydbSetOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "dragonflydb1:6379", setOnlineStorageBenchmarkConstructor)
}
