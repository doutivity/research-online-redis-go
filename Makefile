env-up:
	docker-compose up -d

test:
	docker exec research-online-redis-go-app go test ./... -v -count=1

bench:
	docker exec research-online-redis-go-app go test ./... -v -short -bench=. -benchmem -benchtime=1000000x -count=10 | tee all-bench-10.txt
	benchstat all-bench-10.txt

env-down:
	docker-compose down --remove-orphans -v

pull:
	docker pull redis:latest
	docker pull eqalpha/keydb:latest
	docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest
