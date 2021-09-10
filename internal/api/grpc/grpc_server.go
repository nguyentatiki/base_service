package grpc

import (
	"log"
	"net"

	pb "base_service/internal/api/grpc/proto"
	app "base_service/internal/application"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	handler *app.UserHandler
}

func NewServer(handler *app.UserHandler) *Server {
	return &Server{handler: handler}
}

func Register(port string, server *Server) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
