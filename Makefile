build:
	@go build -o bin/bank_project

run: build
	@./bin/bank_project

test:
	@go test -v ./...