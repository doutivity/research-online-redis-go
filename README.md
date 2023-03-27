# Efficiently store online with Redis and Go
https://dou.ua/forums/topic/35260/

# Testing
```bash
make env-up
make test
make env-down
```
```text
=== RUN   TestRedisHashOnlineStorage
--- PASS: TestRedisHashOnlineStorage (0.01s)
=== RUN   TestKeydbHashOnlineStorage
--- PASS: TestKeydbHashOnlineStorage (0.01s)
=== RUN   TestDragonflydbHashOnlineStorage
--- PASS: TestDragonflydbHashOnlineStorage (0.01s)
PASS
ok  	github.com/doutivity/research-online-redis-go	0.031s
```

# Benchmark
```bash
make bench
# docker exec research-online-redis-go-app go test ./... -v -bench=Redis -benchmem -benchtime=1000000x -count=10 | tee bench-redis-1000000x.txt
# docker exec research-online-redis-go-app go test ./... -v -bench=Keydb -benchmem -benchtime=1000000x -count=10 | tee bench-keydb-1000000x.txt
# benchstat bench-redis-1000000x.txt
# benchstat bench-keydb-1000000x.txt
```text
name                            time/op
RedisHashOnlineStorage-12       18.4µs ±18%
KeydbHashOnlineStorage-12       20.6µs ±27%
RedisSetOnlineStorage-12        19.6µs ±24%
KeydbSetOnlineStorage-12        19.3µs ±27%
RedisSortedSetOnlineStorage-12  21.6µs ±18%
KeydbSortedSetOnlineStorage-12  21.5µs ±27%

name                            alloc/op
RedisHashOnlineStorage-12         240B ± 0%
KeydbHashOnlineStorage-12         240B ± 0%
RedisSetOnlineStorage-12          232B ± 0%
KeydbSetOnlineStorage-12          232B ± 0%
RedisSortedSetOnlineStorage-12    304B ± 0%
KeydbSortedSetOnlineStorage-12    304B ± 0%

name                            allocs/op
RedisHashOnlineStorage-12         10.0 ± 0%
KeydbHashOnlineStorage-12         10.0 ± 0%
RedisSetOnlineStorage-12          10.0 ± 0%
KeydbSetOnlineStorage-12          10.0 ± 0%
RedisSortedSetOnlineStorage-12    10.0 ± 0%
KeydbSortedSetOnlineStorage-12    10.0 ± 0%
```

# Donate
- [Збір для NN бригади від Go-розробника](https://dou.ua/forums/topic/42510/)
