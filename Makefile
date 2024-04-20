build:
	@go build -o bin/kafgo *.go

run: build
	@./bin/kafgo

test:
	@go test -v ./...
