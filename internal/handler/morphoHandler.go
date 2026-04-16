package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type morphoHandler struct{}

func NewMorphoHandler() *morphoHandler {
	return &morphoHandler{}
}

func (h *morphoHandler) GetCurrentVaultPosition(context *gin.Context) {
	context.JSON(http.StatusAccepted, gin.H{
		"data": "test",
	})
}
