build:
	@go build -o bin/bank_project

run: build
	@./bin/bank_project

seed_run: build
	@./bin/bank_project --seed

test:
	@go test -v ./...