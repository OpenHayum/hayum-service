.PHONY: test

run: stop up

mod:
	GO111MODULE=on go mod tidy
	# GO111MODULE=on go mod vendor

up:
	docker-compose -f docker-compose.yaml up --build -d

stop:
	docker-compose stop

down:
	docker-compose down

test:
	docker-compose -f docker-compose.test.yaml up --build --exit-code-from hayum
	docker-compose -f docker-compose.test.yaml down --volumes	