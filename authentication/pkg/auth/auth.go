package auth

import (
	"context"

	"github.com/surajboniwal/connect/authentication/pkg/pb"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{Status: true, Message: "User registered"}, nil
}

func (s *Server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Status: true, Message: "User logged"}, nil
}
