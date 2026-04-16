package main

import (
	"fmt"
	"strconv"

	v1 "github.com/Billy19191/billy191-services-backend/api/v1"
	config "github.com/Billy19191/billy191-services-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	v1.RegisterRoutes(router, config.DB)

	if err := router.Run(":" + strconv.Itoa(config.Port)); err != nil {
		fmt.Println("Error")
	}
}
