package handlers

import (
	"fmt"
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
	// Get the file from the form
	fmt.Println("step 1")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to retrieve file: %v", err)})
		return
	}

	// Call ImportService to process the file
	records := services.ImportService(file)
	fmt.Println("step n-1")
	// Return the parsed records as JSON
	c.JSON(http.StatusOK, records)
}
