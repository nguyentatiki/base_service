package grpc

import (
	"context"
	"log"
	"net"

	pb "base_service/internal/api/grpc/proto"
	"base_service/internal/app"
	"base_service/internal/domain/entities"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// Login implements grpc.Login
func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v", in.GetUsername())
	app.GetUser(&entities.User{Username: in.Username})
	return &pb.LoginResponse{Message: "Hello " + in.GetUsername()}, nil
}

func Register(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
