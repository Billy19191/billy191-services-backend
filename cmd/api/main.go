package main

import (
	"fmt"

	v1 "github.com/Billy19191/billy191-services-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error")
	}
}
