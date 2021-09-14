#!/bin/bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
apt-get update && apt-get install -y protobuf-compiler make
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#for macos
#brew install golang-migrate
export PATH="$PATH:$(go env GOPATH)/bin"
