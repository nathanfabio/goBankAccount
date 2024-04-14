build:
	@go build -o bin/goBankAccount

run: build
	@./bin/goBankAccount

test:
	@go test -v ./...
