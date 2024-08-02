package handlers

import (
	"go-import-manage/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteData(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
