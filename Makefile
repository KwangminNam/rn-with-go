start: build
	@./bin/api

build:
	@go build -o ./bin/api ./cmd/api/main.go

clean:
	@rm -rf ./bin