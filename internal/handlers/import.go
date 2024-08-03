package handlers

import (
	"go-import-manage/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ImportData godoc
// @Summary Import data from an Excel file
// @Description Uploads an Excel file and imports the data
// @Tags import
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /import [post]
func ImportData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}

	if err := services.ImportService(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
}
