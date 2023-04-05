package research_online_redis_go

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

var goStorageConstructor onlineStorageConstructor = func(client *redis.Client) OnlineStorage {
	return NewGoOnlineStorage()
}

func TestGoOnlineStorage(t *testing.T) {
	testOnlineStorage(t, "redis1:6379", goStorageConstructor)
}

func BenchmarkGoOnlineStorage(b *testing.B) {
	benchmarkOnlineStorage(b, "redis1:6379", goStorageConstructor)
}
