package router

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/pkg/auth"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	auth.Router(router.Group("/auth"))

	return router
}
