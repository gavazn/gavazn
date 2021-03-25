deploy: test build-docs build-app

run:
	go run cmd/gavazn/main.go

build-app:
	go build -o ./build/gavazn -i cmd/gavazn/main.go

build-docs:
	apidoc -i ./server/v1 -o ./docs

test:
	go test ./...