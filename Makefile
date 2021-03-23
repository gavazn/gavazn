deploy: test build-app

run:
	go run cmd/gavazn/main.go

build-app:
	go build -o ./build/gavazn -i cmd/gavazn/main.go

test:
	go test ./...