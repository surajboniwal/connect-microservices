package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup, handler *AuthHandler) {

	router.POST("/register", handler.loginHandler)

	router.POST("/login", handler.registerHandler)

}
