package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/internal/pb"
)

type AuthHandler struct {
	conn pb.AuthServiceClient
}

func NewAuthHandler(conn pb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		conn: conn,
	}
}

func (h *AuthHandler) loginHandler(ctx *gin.Context) {

}

func (h *AuthHandler) registerHandler(ctx *gin.Context) {}
