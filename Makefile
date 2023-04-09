build:
	@go build -o bin/go-gorm-restapi

run: build
	@./bin/go-gorm-restapi

test:
	@go test -v ./...
