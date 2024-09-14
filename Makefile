docker_login:
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)

test:
	go test ./...

build:
	docker-compose build

start:
	docker-compose up

build_start: build
	docker-compose up

stop:
	docker-compose down

publish: docker_login
	docker tag shorted $(DOCKER_USERNAME)/shorted:latest
	docker push $(DOCKER_USERNAME)/shorted
