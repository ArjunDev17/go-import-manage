package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondSuccess(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
		Message: msg,
	})
}

func RespondWarning(c *gin.Context, warning string, data interface{}, msg string) {
	c.JSON(http.StatusOK, WarningResponse{
		Success: true,
		Warning: warning,
		Data:    data,
		Message: msg,
	})
}

func RespondError(c *gin.Context, err string, msg string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Success: false,
		Error:   err,
		Message: msg,
	})
}
