# Efficiently store online with Redis and Go
https://dou.ua/forums/topic/35260/

# Support Ukraine 🇺🇦
- [Збір для NN бригади від Go-розробника](https://dou.ua/forums/topic/42510/)
- Help Ukraine via [National Bank of Ukraine](https://bank.gov.ua/en/news/all/natsionalniy-bank-vidkriv-spetsrahunok-dlya-zboru-koshtiv-na-potrebi-armiyi)
- Help Ukraine via [SaveLife](https://savelife.in.ua/en/donate-en/) fund
- More info on [war.ukraine.ua](https://war.ukraine.ua/) and [MFA of Ukraine](https://twitter.com/MFA_Ukraine)

# Databases
| Name                                                    | Stars  | Language                               |
|---------------------------------------------------------|--------|----------------------------------------|
| [Redis](https://github.com/redis/redis)                 | 59100+ | [C](https://dou.ua/forums/tags/C/)     |
| [KeyDB](https://github.com/Snapchat/KeyDB)              | 7100+  | [C++](https://dou.ua/forums/tags/C++/) |
| [DragonflyDB](https://github.com/dragonflydb/dragonfly) | 18300+ | [C++](https://dou.ua/forums/tags/C++/) |

# Data structure usage examples
### Hash
```bash
docker exec research-online-redis-1 redis-cli monitor
```
```bash
docker exec research-online-redis-go-app go test ./... -v -run=TestRedisHashOnlineStorage -count=1
```
```text
flushall
hset "h:online:main" "10000001" "1679800725"
hset "h:online:main" "10000002" "1679800730"
hset "h:online:main" "10000003" "1679800735"
hlen "h:online:main"
rename "h:online:main" "h:online:tmp"
hgetall "h:online:tmp"
```
### Sorted Set
```bash
docker exec research-online-redis-1 redis-cli monitor
```
```bash
docker exec research-online-redis-go-app go test ./... -v -run=TestRedisSortedSetOnlineStorage -count=1
```
```text
flushall
zadd "z:online:main" "1679800725" "10000001"
zadd "z:online:main" "1679800730" "10000002"
zadd "z:online:main" "1679800735" "10000003"
zcard "z:online:main"
rename "z:online:main" "z:online:tmp"
zrange "z:online:tmp" "0" "-1" "withscores"
```
### Set
```bash
docker exec research-online-redis-1 redis-cli monitor
```
```bash
docker exec research-online-redis-go-app go test ./... -v -run=TestRedisSetOnlineStorage -count=1
```
```text
flushall
sadd "s:online:main:1679800725" "10000001"
sadd "s:online:main:1679800730" "10000002"
sadd "s:online:main:1679800735" "10000003"
keys "s:online:main:*"
scard "s:online:main:1679800730"
scard "s:online:main:1679800725"
scard "s:online:main:1679800735"
keys "s:online:main:*"
rename "s:online:main:1679800730" "s:online:tmp"
smembers "s:online:tmp"
rename "s:online:main:1679800725" "s:online:tmp"
smembers "s:online:tmp"
rename "s:online:main:1679800735" "s:online:tmp"
smembers "s:online:tmp"
```

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
--- PASS: TestDragonflydbHashOnlineStorage (0.02s)
=== RUN   TestRedisSetOnlineStorage
--- PASS: TestRedisSetOnlineStorage (0.01s)
=== RUN   TestKeydbSetOnlineStorage
--- PASS: TestKeydbSetOnlineStorage (0.01s)
=== RUN   TestDragonflydbSetOnlineStorage
--- PASS: TestDragonflydbSetOnlineStorage (0.09s)
=== RUN   TestRedisSortedSetOnlineStorage
--- PASS: TestRedisSortedSetOnlineStorage (0.01s)
=== RUN   TestKeydbSortedSetOnlineStorage
--- PASS: TestKeydbSortedSetOnlineStorage (0.01s)
=== RUN   TestDragonflydbSortedSetOnlineStorage
--- PASS: TestDragonflydbSortedSetOnlineStorage (0.05s)
PASS
ok  	github.com/doutivity/research-online-redis-go	0.214s
```

# Benchmark
```bash
make bench
# MODE=sequence go test ./... -v -bench='Redis(Hash|SortedSet|Set)'  -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-redis-1000000x-sequence.txt
# MODE=sequence go test ./... -v -bench='Keydb(Hash|SortedSet|Set)'  -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-keydb-1000000x-sequence.txt
# MODE=sequence go test ./... -v -bench='Dragonflydb(SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-dragonflydb-1000000x-sequence.txt
# MODE=parallel go test ./... -v -bench='Redis(Hash|SortedSet|Set)'  -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-redis-1000000x-parallel.txt
# MODE=parallel go test ./... -v -bench='Keydb(Hash|SortedSet|Set)'  -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-keydb-1000000x-parallel.txt
# MODE=parallel go test ./... -v -bench='Dragonflydb(SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-dragonflydb-1000000x-parallel.txt
# benchstat ./output/bench-redis-1000000x-sequence.txt
# benchstat ./output/bench-keydb-1000000x-sequence.txt
# benchstat ./output/bench-dragonflydb-1000000x-sequence.txt
# benchstat ./output/bench-redis-1000000x-parallel.txt
# benchstat ./output/bench-keydb-1000000x-parallel.txt
# benchstat ./output/bench-dragonflydb-1000000x-parallel.txt
```
| Database name | Data structure | sequence time/op | parallel time/op |
|---------------|----------------|------------------|------------------|
| Redis         | Hash           | 38.8µs ± 9%      | 14.7µs ± 8%      |
| KeyDB         | Hash           | 46.4µs ± 7%      | 15.5µs ± 4%      |
| DragonflyDB   | Hash           | 61.5µs ±10%      | 18.5µs ± 4%      |
| Redis         | Sorted Set     | 42.1µs ±13%      | 17.1µs ±14%      |
| KeyDB         | Sorted Set     | 50.0µs ± 7%      | 16.2µs ± 5%      |
| DragonflyDB   | Sorted Set     | 58.5µs ±13%      | 18.8µs ± 8%      |
| Redis         | Set            | 39.7µs ±13%      | 15.4µs ±12%      |
| KeyDB         | Set            | 47.6µs ± 6%      | 18.3µs ± 6%      |
| DragonflyDB   | Set            | 65.7µs ±13%      | 27.2µs ± 6%      |

# Used memory
```bash
make bench-redis-memory-1m
make bench-keydb-memory-1m
make bench-dragonflydb-memory-1m
# ...
make bench-redis-memory-25m
make bench-keydb-memory-25m
make bench-dragonflydb-memory-25m
```
```bash
# docker exec research-online-redis-1 redis-cli flushall
# docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$^ -bench='Redis(Hash)' -benchmem -benchtime=25000000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-hash-25m.txt
#
# docker exec research-online-redis-1 redis-cli flushall
# docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$^ -bench='Redis(SortedSet)' -benchmem -benchtime=25000000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-sorted-set-25m.txt
#
# docker exec research-online-redis-1 redis-cli flushall
# docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$^ -bench='Redis(Set)' -benchmem -benchtime=25000000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-set-25m.txt
```
| Database name | Data structure | Users      | Memory                                                          |
|---------------|----------------|------------|-----------------------------------------------------------------|
| Redis         | Hash           | 1 000 000  | 62.64 MB                                                        |
| KeyDB         | Hash           | 1 000 000  | 63.49 MB                                                        |
| DragonflyDB   | Hash           | 1 000 000  | 61.51 MB                                                        |
| Redis         | Hash           | 10 000 000 | 727.20 MB                                                       |
| KeyDB         | Hash           | 10 000 000 | 728.14 MB                                                       |
| DragonflyDB   | Hash           | 10 000 000 | 622.59 MB                                                       |
| Redis         | Hash           | 25 000 000 | 1592.14 MB                                                      |
| KeyDB         | Hash           | 25 000 000 | 1593.27 MB                                                      |
| DragonflyDB   | Hash           | 25 000 000 | 1481.70  MB                                                     |
| Redis         | Sorted Set     | 1 000 000  | 91.09 MB                                                        |
| KeyDB         | Sorted Set     | 1 000 000  | 91.93 MB                                                        |
| DragonflyDB   | Sorted Set     | 1 000 000  | 107.87 MB                                                       |
| Redis         | Sorted Set     | 10 000 000 | 1011.78 MB                                                      |
| KeyDB         | Sorted Set     | 10 000 000 | 1012.64 MB                                                      |
| DragonflyDB   | Sorted Set     | 10 000 000 | 1161.64 MB                                                      |
| Redis         | Sorted Set     | 25 000 000 | 2303.58 MB                                                      |
| KeyDB         | Sorted Set     | 25 000 000 | unknown, cause store less then expected, 12583631 from 25000000 |
| DragonflyDB   | Sorted Set     | 25 000 000 | 2675.25 MB                                                      |
| Redis         | Set            | 1 000 000  | 48.14 MB                                                        |
| KeyDB         | Set            | 1 000 000  | 49.02 MB                                                        |
| DragonflyDB   | Set            | 1 000 000  | 32.60 MB                                                        |
| Redis         | Set            | 10 000 000 | 469.57 MB                                                       |
| KeyDB         | Set            | 10 000 000 | 471.44 MB                                                       |
| DragonflyDB   | Set            | 10 000 000 | 297.01 MB                                                       |
| Redis         | Set            | 25 000 000 | 1169.33 MB                                                      |
| KeyDB         | Set            | 25 000 000 | 1175.45 MB                                                      |
| DragonflyDB   | Set            | 25 000 000 | unknown, cause store less then expected, 15443800 from 25000000 |

# Batch insert 10k rows x 10k times benchmark
```bash
make bench-redis-memory-10k-batch-10k
make bench-keydb-memory-10k-batch-10k
make bench-dragonflydb-memory-10k-batch-10k
```
```bash
# docker exec research-online-redis-1 redis-cli flushall
# docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Hash)' -benchmem -benchtime=10000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-hash-10k-batch-10k.txt
#
# docker exec research-online-redis-1 redis-cli flushall
# docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(SortedSet)' -benchmem -benchtime=10000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-sorted-set-10k-batch-10k.txt
#
# docker exec research-online-redis-1 redis-cli flushall
# docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Set)' -benchmem -benchtime=10000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-set-10k-batch-10k.txt
#
# docker exec research-online-redis-1 redis-cli flushall
```

| Database name | Data structure | parallel time/op                                                 |
|---------------|----------------|------------------------------------------------------------------|
| Redis         | Hash           | 8232276 ns/op                                                    |
| KeyDB         | Hash           | 21357358 ns/op                                                   |
| DragonflyDB   | Hash           | 6716157 ns/op                                                    |
| Redis         | Sorted Set     | 12016807 ns/op                                                   |
| KeyDB         | Sorted Set     | 15114051 ns/op                                                   |
| DragonflyDB   | Sorted Set     | 9535106 ns/op                                                    |
| Redis         | Set            | 3187424 ns/op                                                    |
| KeyDB         | Set            | 3233770 ns/op                                                    |
| DragonflyDB   | Set            | unknown, cause store less then expected, 15622200 from 100000000 |

| Database name | Data structure | Memory                                                           |
|---------------|----------------|------------------------------------------------------------------|
| Redis         | Hash           | 6.72 GB                                                          |
| KeyDB         | Hash           | 6.22 GB                                                          |
| DragonflyDB   | Hash           | 5.77 GB                                                          |
| Redis         | Sorted Set     | 9.00 GB                                                          |
| KeyDB         | Sorted Set     | 9.00 GB                                                          |
| DragonflyDB   | Sorted Set     | 10.44 GB                                                         |
| Redis         | Set            | 4.58 GB                                                          |
| KeyDB         | Set            | 4.59 GB                                                          |
| DragonflyDB   | Set            | unknown, cause store less then expected, 15622200 from 100000000 |

# Star history of Redis vs KeyDB vs DragonflyDB
[![Star History Chart](https://api.star-history.com/svg?repos=redis/redis,Snapchat/KeyDB,dragonflydb/dragonfly&type=Date)](https://star-history.com/#redis/redis&Snapchat/KeyDB&dragonflydb/dragonfly&Date)
