build:
	go get github.com/tools/godep
	godep restore
	go build
docker-up: build
	docker-compose down -v --remove-orphans && docker-compose up --force-recreate -d --build
