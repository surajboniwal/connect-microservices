package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/pkg/config"
)

type App struct {
	config *config.Config
	router *gin.Engine
}

func main() {
	config := config.LoadConfig("dev")
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"status": true,
		})
	})

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
