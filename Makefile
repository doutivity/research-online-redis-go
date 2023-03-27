env-up:
	docker-compose up -d

test:
	docker exec research-online-redis-go-app go test ./... -v -count=1

bench-redis:
	docker exec research-online-redis-go-app go test ./... -v -bench=Redis -benchmem -benchtime=1000000x -count=10 | tee bench-redis-1000000x.txt

bench-keydb:
	docker exec research-online-redis-go-app go test ./... -v -bench=Keydb -benchmem -benchtime=1000000x -count=10 | tee bench-keydb-1000000x.txt

bench-dragonflydb:
	docker exec research-online-redis-go-app go test ./... -v -bench=Dragonflydb -benchmem -benchtime=1000000x -count=10 | tee bench-dragonflydb-1000000x.txt

bench: bench-redis bench-keydb # bench-dragonflydb
	benchstat bench-redis-1000000x.txt
	benchstat bench-keydb-1000000x.txt
	# benchstat bench-dragonflydb-1000000x.txt

bench-redis-memory-1m:
	sudo docker exec research-online-redis-1 redis-cli flushall
	sudo docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench=RedisHash -benchmem -benchtime=1000000x -count=1
	sudo docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-hash-1m.txt

	sudo docker exec research-online-redis-1 redis-cli flushall
	sudo docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench=RedisSortedSet -benchmem -benchtime=1000000x -count=1
	sudo docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-sorted-set-1m.txt

	sudo docker exec research-online-redis-1 redis-cli flushall
	sudo docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench=RedisSet -benchmem -benchtime=1000000x -count=1
	sudo docker exec research-online-redis-1 redis-cli info memory | tee redis-memory-set-1m.txt

	sudo docker exec research-online-redis-1 redis-cli flushall

bench-keydb-memory-1m:
	sudo docker exec research-online-keydb-1 keydb-cli flushall
	sudo docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench=KeydbHash -benchmem -benchtime=1000000x -count=1
	sudo docker exec research-online-keydb-1 keydb-cli info memory | tee keydb-memory-hash-1m.txt

	sudo docker exec research-online-keydb-1 keydb-cli flushall
	sudo docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench=KeydbSortedSet -benchmem -benchtime=1000000x -count=1
	sudo docker exec research-online-keydb-1 keydb-cli info memory | tee keydb-memory-sorted-set-1m.txt

	sudo docker exec research-online-keydb-1 keydb-cli flushall
	sudo docker exec research-online-redis-go-app go test ./... -v -run=$$^ -bench=KeydbSet -benchmem -benchtime=1000000x -count=1
	sudo docker exec research-online-keydb-1 keydb-cli info memory | tee keydb-memory-set-1m.txt

	sudo docker exec research-online-keydb-1 keydb-cli flushall

env-down:
	docker-compose down --remove-orphans -v

pull:
	docker pull redis:latest
	docker pull eqalpha/keydb:latest
	docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest
