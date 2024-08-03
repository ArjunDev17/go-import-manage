package services

import (
	"fmt"
	"go-import-manage/internal/models"
	"go-import-manage/internal/utils"
	"mime/multipart"
)

// ImportService handles the import process by calling the ParseExcel function
func ImportService(file *multipart.FileHeader) ([]models.Record, error) {
	// Open the file
	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Call ParseExcel with the file reader
	records, err := utils.ParseExcel(f)
	if err != nil {
		return nil, fmt.Errorf("failed to parse excel file: %w", err)
	}

	return records, nil
}
