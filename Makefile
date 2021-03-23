deploy: test build

run:
	go run cmd/gavazn/main.go

build:
	go build -o ./build/gavazn -i cmd/gavazn/main.go

test:
	go test ./...