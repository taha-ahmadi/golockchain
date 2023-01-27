build:
	@go build -o bin/golackchain

run: build
	@./bin/main

test:
	@go run test -v ./...

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: proto