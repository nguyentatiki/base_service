package api

import (
	"base_service/internal/api/grpc"
	"base_service/internal/api/http"
)

type ApiContainer struct {
	HttpServer *http.Server
	GrpcServer *grpc.Server
}

func NewApiContainer(http *http.Server, grpc *grpc.Server) *ApiContainer {
	return &ApiContainer{HttpServer: http, GrpcServer: grpc}
}
