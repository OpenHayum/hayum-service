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
test-integration:
	GO_ENV=integration go test ./test...

test-integration-coverage:
	GO_ENV=integration go test ./test... -cover -coverprofile=coverage.out
	go tool cover -func=coverage.out

ui-build:
	cd frontend; yarn build

ui-start:
	cd frontend; yarn start