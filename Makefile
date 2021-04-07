PWD = $(shell pwd -L)
IMAGE_NAME = transactions_app

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

logs:
	docker-compose logs -f app

build:
	docker-compose build

enter:
	docker-compose exec app bash