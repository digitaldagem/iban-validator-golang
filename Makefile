SRC_DOCKER_IMAGES := $(shell docker images -q iban-validator-golang-src)
GITHUB_SRC_DOCKER_IMAGES := $(shell docker images -q iban-validator-golang_src)

up:
	docker-compose up -d --build --remove-orphans --timeout 60

up-local:
	docker-compose up --build --remove-orphans

down:
	docker-compose down -v --remove-orphans

	if [ -n "$(SRC_DOCKER_IMAGES)" ]; then docker rmi $(GITHUB_SRC_DOCKER_IMAGES); fi

down-local:
	docker-compose down -v --remove-orphans

	if [ -n "$(SRC_DOCKER_IMAGES)" ]; then docker rmi $(SRC_DOCKER_IMAGES); fi

test:
	go test ./test -v

test-local:
	go test ./test/iban_validator_test.go -v

.PHONY: up up-local down down-local test test-local