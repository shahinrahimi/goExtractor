build:
	@go build -o ./bin/goextractor

run: build
	@./bin/goextractor