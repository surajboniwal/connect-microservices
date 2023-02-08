package auth

import (
	"context"
	"net/http"

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

func (h *AuthHandler) registerHandler(ctx *gin.Context) {
	var body RegisterParams

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.conn.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": response.Message})
}
