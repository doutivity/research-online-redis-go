package research_online_redis_go

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func requireUserOnlinePairsEqual(t testing.TB, expected, actual []UserOnlinePair) {
	t.Helper()

	require.Equal(t, len(expected), len(actual))

	sort.Sort(UserOnlinePairs(expected))
	sort.Sort(UserOnlinePairs(actual))

	require.Equal(t, expected, actual)
}
