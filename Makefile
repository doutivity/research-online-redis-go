env-up:
	docker-compose up -d

test:
	docker exec research-online-redis-go-app go test ./... -v -count=1

bench-redis:
	docker exec research-online-redis-go-app go test ./... -v -bench='Redis(Hash|SortedSet|Set)' -benchmem -benchtime=1000000x -count=10 | tee bench-redis-1000000x.txt

bench-keydb:
	docker exec research-online-redis-go-app go test ./... -v -bench='Keydb(Hash|SortedSet|Set)' -benchmem -benchtime=1000000x -count=10 | tee bench-keydb-1000000x.txt

bench-dragonflydb:
	docker exec research-online-redis-go-app go test ./... -v -bench='Dragonflydb(SortedSet|Set)' -benchmem -benchtime=1000000x -count=10 | tee bench-dragonflydb-1000000x.txt

bench: bench-redis bench-keydb bench-dragonflydb
	benchstat bench-redis-1000000x.txt
	benchstat bench-keydb-1000000x.txt
	benchstat bench-dragonflydb-1000000x.txt

bench-redis-memory-25m:
	docker exec research-online-redis-1 redis-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Hash)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-hash-25m.txt

	docker exec research-online-redis-1 redis-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(SortedSet)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-sorted-set-25m.txt

	docker exec research-online-redis-1 redis-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Set)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-set-25m.txt

	docker exec research-online-redis-1 redis-cli flushall

bench-keydb-memory-25m:
	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(Hash)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee keydb-memory-hash-25m.txt

	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(SortedSet)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee keydb-memory-sorted-set-25m.txt

	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(Set)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee keydb-memory-set-25m.txt

	docker exec research-online-keydb-1 keydb-cli flushall

bench-dragonflydb-memory-25m:
	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(Hash)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee dragonflydb-memory-hash-25m.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(SortedSet)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee dragonflydb-memory-sorted-set-25m.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(Set)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee dragonflydb-memory-set-25m.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall

env-down:
	docker-compose down --remove-orphans -v

pull:
	docker pull redis:latest
	docker pull eqalpha/keydb:latest
	docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest
