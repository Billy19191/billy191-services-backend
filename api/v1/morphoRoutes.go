package v1

import (
	"github.com/Billy19191/billy191-services-backend/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerMorphoRoutes(router *gin.Engine, db *gorm.DB) {
	morphoRoutes := handler.NewMorphoHandler()

	v1 := router.Group("api/v1")

	{
		v1.GET("/vaultPosition", morphoRoutes.GetCurrentVaultPosition)
	}

}
