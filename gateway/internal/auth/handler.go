package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/connect/gateway/internal/pb"
	"github.com/surajboniwal/connect/gateway/internal/util"
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

	if err := ctx.ShouldBind(&body); err != nil {
		e := util.ParseError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "error": e})
		return
	}

	response, err := h.conn.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": "Something went wrong!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": response.Status})
}
