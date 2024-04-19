build:
	@go build -o bin/kafgo main.go

run: build
	@./bin/kafgo

test:
	@go test -v ./...
