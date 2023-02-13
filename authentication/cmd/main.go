package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/surajboniwal/connect/authentication/internal/auth"
	"github.com/surajboniwal/connect/authentication/internal/config"
	"github.com/surajboniwal/connect/authentication/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig("dev")

	if err != nil {
		c = config.Config{
			Port: os.Getenv("PORT"),
		}
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Port))

	if err != nil {
		log.Fatalf("Unable to listen to port %s", c.Port)
	}

	server := grpc.NewServer()

	authServer := auth.Server{}

	pb.RegisterAuthServiceServer(server, &authServer)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %s", err)
	}
}
