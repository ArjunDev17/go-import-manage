package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/import", ImportData)
	router.GET("/view", ViewData)
	router.PUT("/edit/:id", EditData)
	router.DELETE("/delete/:id", DeleteData)
}
