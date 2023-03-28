# Efficiently store online with Redis and Go
https://dou.ua/forums/topic/35260/

# Support Ukraine 🇺🇦
- [Збір для NN бригади від Go-розробника](https://dou.ua/forums/topic/42510/)
- Help Ukraine via [National Bank of Ukraine](https://bank.gov.ua/en/news/all/natsionalniy-bank-vidkriv-spetsrahunok-dlya-zboru-koshtiv-na-potrebi-armiyi)
- Help Ukraine via [SaveLife](https://savelife.in.ua/en/donate-en/) fund
- More info on [war.ukraine.ua](https://war.ukraine.ua/) and [MFA of Ukraine](https://twitter.com/MFA_Ukraine)

# Databases
| Name                                                    | Stars  |
|---------------------------------------------------------|--------|
| [Redis](https://github.com/redis/redis)                 | 59100+ |
| [KeyDB](https://github.com/Snapchat/KeyDB)              | 7100+  |
| [DragonflyDB](https://github.com/dragonflydb/dragonfly) | 18000+ |

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
# go test ./... -v -bench='Redis(Hash|SortedSet|Set)'  -benchmem -benchtime=1000000x -count=10 | tee bench-redis-1000000x.txt
# go test ./... -v -bench='Keydb(Hash|SortedSet|Set)'  -benchmem -benchtime=1000000x -count=10 | tee bench-keydb-1000000x.txt
# go test ./... -v -bench='Dragonflydb(SortedSet|Set)' -benchmem -benchtime=1000000x -count=10 | tee bench-dragonflydb-1000000x.txt
# benchstat bench-redis-1000000x.txt
# benchstat bench-keydb-1000000x.txt
# benchstat bench-dragonflydb-1000000x.txt
```
| Database name | Data structure | time/op                                                                                  |
|---------------|----------------|------------------------------------------------------------------------------------------|
| Redis         | Hash           | 18.4µs ±18%                                                                              |
| KeyDB         | Hash           | 20.6µs ±27%                                                                              |
| DragonflyDB   | Hash           | [unknown, cause server misbehaving](https://github.com/dragonflydb/dragonfly/issues/993) |
| Redis         | Sorted Set     | 21.6µs ±18%                                                                              |
| KeyDB         | Sorted Set     | 21.5µs ±27%                                                                              |
| DragonflyDB   | Sorted Set     | 23.1µs ± 9%                                                                              |
| Redis         | Set            | 19.6µs ±24%                                                                              |
| KeyDB         | Set            | 19.3µs ±27%                                                                              |
| DragonflyDB   | Set            | 17.2µs ±25%                                                                              |

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
# docker exec research-online-redis-go-app go test ./... -v -run=$^ -bench='Redis(Hash)' -benchmem -benchtime=25000000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-hash-25m.txt
#
# docker exec research-online-redis-1 redis-cli flushall
# docker exec research-online-redis-go-app go test ./... -v -run=$^ -bench='Redis(SortedSet)' -benchmem -benchtime=25000000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-sorted-set-25m.txt
#
# docker exec research-online-redis-1 redis-cli flushall
# docker exec research-online-redis-go-app go test ./... -v -run=$^ -bench='Redis(Set)' -benchmem -benchtime=25000000x -count=1
# docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-set-25m.txt
```
| Database name | Data structure | Users      | Memory                                                                                   |
|---------------|----------------|------------|------------------------------------------------------------------------------------------|
| Redis         | Hash           | 1 000 000  | 62.64 MB                                                                                 |
| KeyDB         | Hash           | 1 000 000  | 63.49 MB                                                                                 |
| DragonflyDB   | Hash           | 1 000 000  | [unknown, cause server misbehaving](https://github.com/dragonflydb/dragonfly/issues/993) |
| Redis         | Hash           | 10 000 000 | 727.20 MB                                                                                |
| KeyDB         | Hash           | 10 000 000 | 728.14 MB                                                                                |
| DragonflyDB   | Hash           | 10 000 000 | [unknown, cause server misbehaving](https://github.com/dragonflydb/dragonfly/issues/993) |
| Redis         | Hash           | 25 000 000 | 1592.14 MB                                                                               |
| KeyDB         | Hash           | 25 000 000 | 1593.27 MB                                                                               |
| DragonflyDB   | Hash           | 25 000 000 | [unknown, cause server misbehaving](https://github.com/dragonflydb/dragonfly/issues/993) |
| Redis         | Sorted Set     | 1 000 000  | 91.09 MB                                                                                 |
| KeyDB         | Sorted Set     | 1 000 000  | 91.93 MB                                                                                 |
| DragonflyDB   | Sorted Set     | 1 000 000  | 107.87 MB                                                                                |
| Redis         | Sorted Set     | 10 000 000 | 1011.78 MB                                                                               |
| KeyDB         | Sorted Set     | 10 000 000 | 1012.64 M                                                                                |
| DragonflyDB   | Sorted Set     | 10 000 000 | 1161.64 MB                                                                               |
| Redis         | Sorted Set     | 25 000 000 | 2303.58 MB                                                                               |
| KeyDB         | Sorted Set     | 25 000 000 | unknown, cause store less then expected, 12583631 from 25000000                          |
| DragonflyDB   | Sorted Set     | 25 000 000 | 2675.25 MB                                                                               |
| Redis         | Set            | 1 000 000  | 48.14 MB                                                                                 |
| KeyDB         | Set            | 1 000 000  | 49.02 MB                                                                                 |
| DragonflyDB   | Set            | 1 000 000  | 32.60 MB                                                                                 |
| Redis         | Set            | 10 000 000 | 469.57 MB                                                                                |
| KeyDB         | Set            | 10 000 000 | 471.44 MB                                                                                |
| DragonflyDB   | Set            | 10 000 000 | 297.01 MB                                                                                |
| Redis         | Set            | 25 000 000 | 1169.33 MB                                                                               |
| KeyDB         | Set            | 25 000 000 | 1175.45 MB                                                                               |
| DragonflyDB   | Set            | 25 000 000 | unknown, cause store less then expected, 15443800 from 25000000                          |

# Star history of Redis vs KeyDB vs DragonflyDB
[![Star History Chart](https://api.star-history.com/svg?repos=redis/redis,Snapchat/KeyDB,dragonflydb/dragonfly&type=Date)](https://star-history.com/#redis/redis&Snapchat/KeyDB&dragonflydb/dragonfly&Date)
