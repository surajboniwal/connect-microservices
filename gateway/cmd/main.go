package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/pkg/config"
	"github.com/surajboniwal/connect/gateway/pkg/router"
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

	app.StartServer()
}

func (app *App) StartServer() {
	err := app.router.Run(app.config.ServerAddress)

	if err != nil {
		log.Panic("Unable to start server", err)
	}
}
