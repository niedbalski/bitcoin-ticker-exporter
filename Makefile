up:
	docker-compose down -v --remove-orphans && docker-compose up --force-recreate -d --build

build:
	docker-compose build

push:
	docker push us.gcr.io/lincl-206618/bitcoin-ticker-exporter
