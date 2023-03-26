package research_online_redis_go

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

var hashOnlineStorageConstructor onlineStorageConstructor = func(client *redis.Client) OnlineStorage {
	return NewHashOnlineStorage(client)
}

func TestRedisHashOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "redis1:6379", hashOnlineStorageConstructor)
}

func TestKeydbHashOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "keydb1:6379", hashOnlineStorageConstructor)
}

func TestDragonflydbHashOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "dragonflydb1:6379", hashOnlineStorageConstructor)
}

func BenchmarkRedisHashOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "redis1:6379", hashOnlineStorageConstructor)
}

func BenchmarkKeydbHashOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "keydb1:6379", hashOnlineStorageConstructor)
}

func BenchmarkDragonflydbHashOnlineStorage(b *testing.B) {
	if testing.Short() {
		return
	}

	benchmarkOnlineStorage(b, "dragonflydb1:6379", hashOnlineStorageConstructor)
}
