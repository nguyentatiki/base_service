//go:build wireinject
// +build wireinject

package internal

import (
	"base_service/internal/api"
	"base_service/internal/api/grpc"
	"base_service/internal/api/http"
	app "base_service/internal/application"
	"base_service/internal/infrastructure/persistent"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var container = wire.NewSet(
	api.NewApiContainer,
)

var apiSet = wire.NewSet(
	grpc.NewServer,
	http.NewServer,
)

var reposet = wire.NewSet(
	persistent.NewUserRepository,
)

var handlerSet = wire.NewSet(
	app.NewUserHandler,
)

func InitializeContainer(db *sqlx.DB) *api.ApiContainer {
	wire.Build(handlerSet, reposet, apiSet, container)
	return &api.ApiContainer{}
}
