package main

import (
	"fmt"
	"log"
	"net"

	"github.com/surajboniwal/connect/authentication/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	config := config.LoadConfig("dev")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))

	if err != nil {
		log.Fatalf("Unable to listen to port %s", config.Port)
	}

	server := grpc.NewServer()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %s", err)
	}

}
