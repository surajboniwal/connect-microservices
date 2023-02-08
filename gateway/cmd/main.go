package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/internal/auth"
	"github.com/surajboniwal/connect/gateway/internal/config"
	"github.com/surajboniwal/connect/gateway/internal/pb"
	"github.com/surajboniwal/connect/gateway/internal/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	config *config.Config
	router *gin.Engine
}

func main() {
	config := config.LoadConfig("dev")

	authClient := connectToAuthService(&config)

	handlers := router.Handlers{
		AuthHandler: auth.NewAuthHandler(authClient),
	}

	router := router.NewRouter(&handlers)

	app := App{
		config: &config,
		router: router,
	}

	app.StartServer()
}

func connectToAuthService(config *config.Config) pb.AuthServiceClient {
	conn, err := grpc.Dial(config.AuthServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect to auth service %v", err)
	}

	client := pb.NewAuthServiceClient(conn)

	return client
}

func (app *App) StartServer() {
	err := app.router.Run(app.config.ServerAddress)

	if err != nil {
		log.Panic("Unable to start server", err)
	}
}
