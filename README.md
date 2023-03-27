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
```
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

# Used memory
### Redis Hash
```bash
go test ./... -v -run=$^ -bench=RedisHash -benchmem -benchtime=1000000x -count=1
redis-cli info memory
```
```text
used_memory_human:      62.64M
used_memory_rss_human:  75.66M
used_memory_peak_human: 91.32M
```
### KeyDB Hash
```bash
go test ./... -v -run=$^ -bench=KeydbHash -benchmem -benchtime=1000000x -count=1
keydb-cli info memory
```
```text
used_memory_human:      63.49M
used_memory_rss_human:  84.80M
used_memory_peak_human: 92.60M
```
### Redis Sorted Set
```bash
go test ./... -v -run=$^ -bench=RedisSortedSet -benchmem -benchtime=1000000x -count=1
redis-cli info memory
```
```text
used_memory_human:      91.09M
used_memory_rss_human:  99.88M
used_memory_peak_human: 91.32M
```
### KeyDB Sorted Set
```bash
go test ./... -v -run=$^ -bench=KeydbSortedSet -benchmem -benchtime=1000000x -count=1
keydb-cli info memory
```
```text
used_memory_human:      91.93M
used_memory_rss_human:  115.38M
used_memory_peak_human: 92.60M
```
### Redis Set
```bash
go test ./... -v -run=$^ -bench=RedisSet -benchmem -benchtime=1000000x -count=1
redis-cli info memory
```
```text
used_memory_human:      48.14M
used_memory_rss_human:  59.30M
used_memory_peak_human: 91.32M
```
### KeyDB Set
```bash
go test ./... -v -run=$^ -bench=KeydbSet -benchmem -benchtime=1000000x -count=1
keydb-cli info memory
```
```text
used_memory_human:      49.02M
used_memory_rss_human:  70.49M
used_memory_peak_human: 92.60M
```

# Donate
- [Збір для NN бригади від Go-розробника](https://dou.ua/forums/topic/42510/)
