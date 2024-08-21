package handlers

import (
	middleware "go-import-manage/internal/midleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	// Grouping the routes under "/api"
	dataGroup := router.Group("/api")
	{

		dataGroup.POST("/import", middleware.FileValidationMiddleware(), ImportData)

		dataGroup.GET("/view", ViewData)
		dataGroup.PUT("/edit/:id", EditData)
		dataGroup.DELETE("/delete/:id", DeleteData)
	}
}
