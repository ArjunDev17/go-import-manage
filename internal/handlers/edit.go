package handlers

import (
	"go-import-manage/internal/models"
	"go-import-manage/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EditData godoc
// @Summary Edit a specific record
// @Description Edits a specific record
// @Tags edit
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Param record body models.Record true "Record"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /edit/{id} [put]
func EditData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

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
