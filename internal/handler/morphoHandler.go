package handler

import (
	"net/http"
	"strconv"

	service "github.com/Billy19191/billy191-services-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type morphoHandler struct {
	service *service.MorphoService
}

func NewMorphoHandler(service *service.MorphoService) *morphoHandler {
	return &morphoHandler{
		service: service,
	}
}

func (h *morphoHandler) GetCurrentVaultPosition(context *gin.Context) {
	walletAddress := context.Query("walletAddress")
	chainID := context.Query("chainID")

	chainIDInt, err := strconv.Atoi(chainID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid chainID",
		})
		return
	}

	result, err := h.service.GetVaultPositionByWallet(walletAddress, chainIDInt)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": result.Data.Vault,
	})
}
