package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	router.POST("/register", loginHandler)

	router.POST("/login", registerHandler)

}
