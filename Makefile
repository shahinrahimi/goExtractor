build:
	@go build -o ./bin/goextractor && echo "Build successful!"

run: build
	@./bin/goextractor