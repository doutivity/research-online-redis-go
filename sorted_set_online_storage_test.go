package research_online_redis_go

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

var sortedSetOnlineStorageConstructor onlineStorageConstructor = func(client *redis.Client) OnlineStorage {
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
