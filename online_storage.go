package research_online_redis_go

import (
	"context"
)

type OnlineStorage interface {
	Store(ctx context.Context, pair UserOnlinePair) error
	Count(ctx context.Context) (int64, error)
	GetAndClear(ctx context.Context) ([]UserOnlinePair, error)
}
