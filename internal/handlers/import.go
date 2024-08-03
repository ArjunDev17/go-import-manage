package handlers

import (
	"fmt"
	"go-import-manage/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ImportData handles the import request in the Gin framework
func ImportData(c *gin.Context) {
	// Get the file from the form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to retrieve file: %v", err)})
		return
	}

	// Call ImportService to process the file
	records, err := services.ImportService(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to import file: %v", err)})
		return
	}

	// Return the parsed records as JSON
	c.JSON(http.StatusOK, records)
}
