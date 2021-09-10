//go:build wireinject

package internal

import (
	"base_service/internal/api/grpc"
	app "base_service/internal/application"
	"base_service/internal/infrastructure/persistent"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var ApiSet = wire.NewSet(
	grpc.NewServer,
)

var Reposet = wire.NewSet(
	persistent.NewUserRepository,
)

var HandlerSet = wire.NewSet(
	app.NewUserHandler,
)

func InitializeContainer(db *sqlx.DB) *grpc.Server {
	wire.Build(HandlerSet, Reposet, ApiSet)
	return &grpc.Server{}
}
