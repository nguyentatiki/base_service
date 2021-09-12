include .env

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/api/grpc/proto/user.proto

migrate-up:
	./base_service migrate up

migrate-down:
	./base_service migrate down

create-migration-files:
	migrate create -ext sql -dir migrations ${filename}

install:
	sh install.sh
