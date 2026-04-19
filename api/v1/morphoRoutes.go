package v1

import (
	"github.com/Billy19191/billy191-services-backend/internal/client"
	"github.com/Billy19191/billy191-services-backend/internal/handler"
	"github.com/Billy19191/billy191-services-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerMorphoRoutes(router *gin.Engine, db *gorm.DB) {
	morphoClient := client.NewMorphoClient()
	morphoService := service.NewMorphoService(db, morphoClient)
	morphoHandler := handler.NewMorphoHandler(morphoService)

	v1 := router.Group("/api/v1")

	{
		v1.GET("/vaultPosition", morphoHandler.GetCurrentVaultPosition)
	}

}
