package handlers

import (
	"go-import-manage/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewData(c *gin.Context) {
	data, err := services.ViewService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
