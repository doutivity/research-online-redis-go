package research_online_redis_go

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Client(ctx context.Context, addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr, // "localhost:6379",
		Password: "",   // no password set
		DB:       0,    // use default DB
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
