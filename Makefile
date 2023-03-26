env-up:
	docker-compose up -d

test:
	docker exec research-online-redis-go-app go test ./... -v -count=1

env-down:
	docker-compose down

pull:
	docker pull redis:latest
	docker pull eqalpha/keydb:latest
	docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest
