.PHONY: test

run: stop up

mod:
	# This make rule requires Go 1.11+
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

up:
	docker-compose -f docker-compose.yaml up --build

stop:
	docker-compose stop

down:
	docker-compose down

test:
	docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yaml down --volumes	