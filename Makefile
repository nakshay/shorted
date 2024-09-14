
test:
	go test ./...

build:
	docker-compose build

start:
	docker-compose up

stop:
	docker-compose down

