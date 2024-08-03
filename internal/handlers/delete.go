package handlers

import (
	"go-import-manage/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteData godoc
// @Summary Delete a specific record
// @Description Deletes a specific record from the database and cache
// @Tags delete
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /delete/{id} [delete]
func DeleteData(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
