package v1

import (
	"github.com/Billy19191/billy191-services-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.Engine) {
	userHandler := handler.NewUserHandler()
	v1 := router.Group("api/v1")

	{
		v1.GET("/user", userHandler.GetUser)
	}
}
