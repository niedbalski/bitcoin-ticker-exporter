up:
	docker-compose down -v --remove-orphans && docker-compose up --force-recreate -d --build
