build:
	@go build -o bin/ristorante
run: build
	@./bin/ristorante
test:
	@go test -v ./...