package v1

import (
	"github.com/Billy19191/billy191-services-backend/internal/client"
	"github.com/Billy19191/billy191-services-backend/internal/handler"
	"github.com/Billy19191/billy191-services-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerAccountableRoutes(router *gin.Engine, db *gorm.DB) {
	accountableClient := client.NewAccountableClient()
	accountableService := service.NewAccountableService(db, accountableClient)
	accountableHandler := handler.NewAccountableHandler(accountableService)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/position-accountable", accountableHandler.GetPositionAccountableData)
	}

}
