package middleware

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func FileValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the file from the form

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			c.Abort() // Stop the request from proceeding
			return
		}

		// Get the file extension
		ext := strings.ToLower(filepath.Ext(file.Filename))

		// Allow only .csv and .xlsx files
		if ext != ".csv" && ext != ".xlsx" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported file type. Only .csv and .xlsx are allowed."})
			c.Abort() // Stop the request from proceeding
			return
		}

		// Proceed to the next handler if the file is valid
		c.Next()
	}
}
func validToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
