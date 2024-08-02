package handlers

import (
	"go-import-manage/internal/models"
	"go-import-manage/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditData(c *gin.Context) {
	id := c.Param("id")

	var record models.Record

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.EditService(id, record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
