package grpc

import (
	"context"
	"log"
	"net"
	"os"

	pb "base_service/internal/api/grpc/proto"
	"base_service/internal/app"
	"base_service/internal/domain/entities"
	"base_service/internal/infrastructure/persistent"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	db *sqlx.DB
}

// Login implements grpc.Login
func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: %v", in.GetUsername())
	repo := persistent.NewUserRepository(s.db)
	user, err := app.NewUserHandler(repo).GetUser(&entities.User{Username: in.Username})
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{Username: user.Username, Email: user.Email, Phone: user.PhoneNumber}, nil
}

func Register(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	conn := os.Getenv("MYSQL_CONNECTION")
	db, err := sqlx.Connect("mysql", conn)

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
