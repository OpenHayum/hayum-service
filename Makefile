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

# use this for local integration test
# using docker-compose is a bit slow
integration_test:
	GO_ENV=integration go test ./... -tags=integration

test_coverage:
	GO_ENV=integration go test ./... -tags=integration -cover -coverprofile=coverage.out
	go tool cover -func=coverage.out