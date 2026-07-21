package handler

import (
	"net/http"
	"strconv"

	"github.com/Billy19191/billy191-services-backend/internal/model"
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

func (h *morphoHandler) GetCurrentDepositVaultPosition(context *gin.Context) {
	walletAddress := context.Query("walletAddress")
	chainID := context.Query("chainID")

	if walletAddress == "" {
		badRequest := http.StatusBadRequest
		message := "walletAddress is required"
		context.JSON(http.StatusBadRequest, model.FailedResponseModel{
			ResponseCode: &badRequest,
			Error:        &message,
		})
		return
	}

	chainIDInt, err := strconv.Atoi(chainID)
	if err != nil {
		badRequest := http.StatusBadRequest
		message := "Invalid chainID"
		context.JSON(http.StatusBadRequest, model.FailedResponseModel{
			ResponseCode: &badRequest,
			Error:        &message,
		})
		return
	}

	result, err := h.service.GetVaultPositionByWallet(walletAddress, chainIDInt)
	if err != nil {
		internalError := http.StatusInternalServerError
		message := err.Error()
		context.JSON(http.StatusInternalServerError, model.FailedResponseModel{
			ResponseCode: &internalError,
			Error:        &message,
		})
		return
	}

	successCode := http.StatusOK

	context.JSON(http.StatusOK, model.SuccessResponseModel{
		ResponseCode: &successCode,
		ResponseData: result.Data.Vault,
	})
}

func (h *morphoHandler) GetCurrentBorrowVaultPosition(context *gin.Context) {
	walletAddress := context.Query("walletAddress")
	chainID := context.Query("chainID")

	if walletAddress == "" {
		badRequest := http.StatusBadRequest
		message := "walletAddress is required"
		context.JSON(http.StatusBadRequest, model.FailedResponseModel{
			ResponseCode: &badRequest,
			Error:        &message,
		})
		return
	}

	chainIDInt, err := strconv.Atoi(chainID)
	if err != nil {
		badRequest := http.StatusBadRequest
		message := "Invalid chainID"
		context.JSON(http.StatusBadRequest, model.FailedResponseModel{
			ResponseCode: &badRequest,
			Error:        &message,
		})
		return
	}

	result, err := h.service.GetBorrowPositionByWallet(walletAddress, chainIDInt)
	if err != nil {
		internalError := http.StatusInternalServerError
		message := err.Error()
		context.JSON(http.StatusInternalServerError, model.FailedResponseModel{
			ResponseCode: &internalError,
			Error:        &message,
		})
		return
	}

	successCode := http.StatusOK

	context.JSON(http.StatusOK, model.SuccessResponseModel{
		ResponseCode: &successCode,
		ResponseData: result.Data.Borrow,
	})
}
