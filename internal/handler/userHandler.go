package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string]string{
			"user": "Billy191",
		},
	})
}
