generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/api/grpc/proto/user.proto

migrate-up:
	migrate -path=migrations -database=mysql://"$(MYSQL_CONNECTION)" up

install:
	sh install.sh
