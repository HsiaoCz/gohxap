package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomePageHandler struct{}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{}
}

func (h *HomePageHandler) HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the home page",
	})
}
