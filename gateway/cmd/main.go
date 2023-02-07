package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/pkg/config"
	"github.com/surajboniwal/connect/gateway/pkg/pb"
	"github.com/surajboniwal/connect/gateway/pkg/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	config *config.Config
	router *gin.Engine
}

func main() {
	config := config.LoadConfig("dev")
	router := router.NewRouter()

	app := App{
		config: &config,
		router: router,
	}

	app.connectToAuthService()

	app.StartServer()
}

func (app *App) connectToAuthService() {
	conn, err := grpc.Dial(app.config.AuthServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect to auth service %v", err)
	}

	pb.NewAuthServiceClient(conn)
}

func (app *App) StartServer() {
	err := app.router.Run(app.config.ServerAddress)

	if err != nil {
		log.Panic("Unable to start server", err)
	}
}
