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

env-down:
	docker-compose down --remove-orphans -v

pull:
	docker pull redis:latest
	docker pull eqalpha/keydb:latest
	docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest
