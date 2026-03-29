package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	registerUserRoutes(router)
}
