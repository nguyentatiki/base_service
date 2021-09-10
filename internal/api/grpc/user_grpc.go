package grpc

import (
	pb "base_service/internal/api/grpc/proto"
	"base_service/internal/domain/entities"
	"context"
	"log"
)

// GetUser implements grpc.GetUser
func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: %v", in.GetUsername())
	user, err := s.handler.GetUser(&entities.User{Username: in.Username})
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{Username: user.Username, Email: user.Email, Phone: user.PhoneNumber}, nil
}
