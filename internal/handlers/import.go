package handlers

import (
	"fmt"
	"go-import-manage/internal/services"
	"go-import-manage/internal/utils"

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
	file, err := c.FormFile("file")
	if err != nil {
		utils.RespondError(c, fmt.Sprintf("failed to retrieve file: %v", err), "File retrieval failed")
		return
	}

	// Call ImportService to process the file
	err = services.ImportService(file)
	if err != nil {
		utils.RespondError(c, fmt.Sprintf("failed to import data: %v", err), "Import failed")
		return
	}

	// Return success response
	utils.RespondSuccess(c, nil, "Data imported successfully")
}
