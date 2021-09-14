include .env

.PHONY: generate-proto
generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/api/grpc/proto/user.proto

.PHONY: migrate-up
migrate-up:
	./base_service migrate up

.PHONY: migrate-down
migrate-down:
	./base_service migrate down

.PHONY:create-migration-files
create-migration-files:
	migrate create -ext sql -dir migrations ${filename}

.PHONY: install
install:
	sh install.sh
