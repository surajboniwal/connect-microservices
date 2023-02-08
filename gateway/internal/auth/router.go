package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup, handler *AuthHandler) {

	router.POST("/register", handler.registerHandler)

	router.POST("/login", handler.loginHandler)

}
