package router

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/internal/auth"
)

type Handlers struct {
	AuthHandler *auth.AuthHandler
}

func NewRouter(handlers *Handlers) *gin.Engine {
	router := gin.Default()

	auth.Router(router.Group("/auth"), handlers.AuthHandler)

	return router
}
