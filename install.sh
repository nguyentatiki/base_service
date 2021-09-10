go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go get github.com/google/wire/cmd/wire
#for macos
#brew install golang-migrate
export PATH="$PATH:$(go env GOPATH)/bin"
