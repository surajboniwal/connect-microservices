package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/internal/pb"
	"golang.org/x/net/context"
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
	loginResponse, err := h.conn.Login(context.Background(), &pb.LoginRequest{
		Email:    "surajboniwal18@gmail.com",
		Password: "suraj1335",
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"hello": loginResponse.Message,
	})
}

func (h *AuthHandler) registerHandler(ctx *gin.Context) {}
