env-up:
	docker-compose up -d

test:
	docker exec research-online-redis-go-app go test ./... -v -count=1

bench-redis-sequence:
	docker exec -e MODE=sequence research-online-redis-go-app go test ./... -v -bench='Redis(Hash|SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-redis-1000000x-sequence.txt

bench-keydb-sequence:
	docker exec -e MODE=sequence research-online-redis-go-app go test ./... -v -bench='Keydb(Hash|SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-keydb-1000000x-sequence.txt

bench-dragonflydb-sequence:
	docker exec -e MODE=sequence research-online-redis-go-app go test ./... -v -bench='Dragonflydb(SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-dragonflydb-1000000x-sequence.txt

bench-redis-parallel:
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -bench='Redis(Hash|SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-redis-1000000x-parallel.txt

bench-keydb-parallel:
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -bench='Keydb(Hash|SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-keydb-1000000x-parallel.txt

bench-dragonflydb-parallel:
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -bench='Dragonflydb(SortedSet|Set)' -benchmem -benchtime=1000000x -count=5 | tee ./output/bench-dragonflydb-1000000x-parallel.txt

bench: bench-redis-sequence bench-keydb-sequence bench-dragonflydb-sequence bench-redis-parallel bench-keydb-parallel bench-dragonflydb-parallel
	benchstat bench-redis-1000000x-sequence.txt
	benchstat bench-keydb-1000000x-sequence.txt
	benchstat bench-dragonflydb-1000000x-sequence.txt
	benchstat bench-redis-1000000x-parallel.txt
	benchstat bench-keydb-1000000x-parallel.txt
	benchstat bench-dragonflydb-1000000x-parallel.txt

bench-redis-memory-25m:
	docker exec research-online-redis-1 redis-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Hash)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-hash-25m.txt

	docker exec research-online-redis-1 redis-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(SortedSet)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-sorted-set-25m.txt

	docker exec research-online-redis-1 redis-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Set)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-set-25m.txt

	docker exec research-online-redis-1 redis-cli flushall

bench-keydb-memory-25m:
	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(Hash)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee ./output/keydb-memory-hash-25m.txt

	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(SortedSet)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee ./output/keydb-memory-sorted-set-25m.txt

	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(Set)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee ./output/keydb-memory-set-25m.txt

	docker exec research-online-keydb-1 keydb-cli flushall

bench-dragonflydb-memory-25m:
	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(Hash)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee ./output/dragonflydb-memory-hash-25m.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(SortedSet)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee ./output/dragonflydb-memory-sorted-set-25m.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec -e MODE=parallel research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(Set)' -benchmem -benchtime=25000000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee ./output/dragonflydb-memory-set-25m.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall

bench-redis-memory-10k-batch-10k:
	docker exec research-online-redis-1 redis-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Hash)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-hash-10k-batch-10k.txt

	docker exec research-online-redis-1 redis-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(SortedSet)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-sorted-set-10k-batch-10k.txt

	docker exec research-online-redis-1 redis-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Redis(Set)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-redis-1 redis-cli info memory | tee ./output/redis-memory-set-10k-batch-10k.txt

	docker exec research-online-redis-1 redis-cli flushall

bench-keydb-memory-10k-batch-10k:
	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(Hash)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee ./output/keydb-memory-hash-10k-batch-10k.txt

	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(SortedSet)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee ./output/keydb-memory-sorted-set-10k-batch-10k.txt

	docker exec research-online-keydb-1 keydb-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Keydb(Set)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-keydb-1 keydb-cli info memory | tee ./output/keydb-memory-set-10k-batch-10k.txt

	docker exec research-online-keydb-1 keydb-cli flushall

bench-dragonflydb-memory-10k-batch-10k:
	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(Hash)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee ./output/dragonflydb-memory-hash-10k-batch-10k.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(SortedSet)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee ./output/dragonflydb-memory-sorted-set-10k-batch-10k.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall
	docker exec -e MODE=parallel -e BATCH=10000 research-online-redis-go-app go test ./... -v -run=$$^ -bench='Dragonflydb(Set)' -benchmem -benchtime=10000x -count=1
	docker exec research-online-dragonflydb-1 redis-cli info memory | tee ./output/dragonflydb-memory-set-10k-batch-10k.txt

	docker exec research-online-dragonflydb-1 redis-cli flushall

env-down:
	docker-compose down --remove-orphans -v

pull:
	docker pull redis:latest
	docker pull eqalpha/keydb:latest
	docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest
