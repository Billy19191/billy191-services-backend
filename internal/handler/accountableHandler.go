package handler

import (
	"github.com/Billy19191/billy191-services-backend/internal/model"
	service "github.com/Billy19191/billy191-services-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AccountableHandler struct {
	service *service.AccountableService
}

func NewAccountableHandler(service *service.AccountableService) *AccountableHandler {
	return &AccountableHandler{
		service: service,
	}
}

func (h *AccountableHandler) GetPositionAccountableData(context *gin.Context) {
	walletAddress := context.Query("walletAddress")

	if walletAddress == "" {
		context.JSON(400, gin.H{
			"error": "walletAddress is required",
		})
		return
	}

	res, err := h.service.GetPositionAccountableData(walletAddress)

	if err != nil {
		errMsg := err.Error()
		context.JSON(500, model.FailedResponseModel{
			ResponseCode: &[]int{500}[0],
			Error:        &errMsg,
		})
		return
	}

	context.JSON(200, model.SuccessResponseModel{
		ResponseCode: &[]int{200}[0],
		ResponseData: res,
	})
}
