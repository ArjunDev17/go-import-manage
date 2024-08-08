package handlers

import (
	"go-import-manage/internal/services"
	"go-import-manage/internal/utils"

	"github.com/gin-gonic/gin"
)

// ViewData godoc
// @Summary View imported data
// @Description Retrieves the imported data from Redis cache or MySQL database
// @Tags view
// @Produce json
// @Success 200 {array} models.Record
// @Failure 500 {object} map[string]string
// @Router /view [get]
func ViewData(c *gin.Context) {
	data, err := services.ViewService()
	if err != nil {
		utils.RespondError(c, err.Error(), "Failed to retrieve data")
		return
	}
	utils.RespondSuccess(c, data, "Data retrieved successfully")
}
