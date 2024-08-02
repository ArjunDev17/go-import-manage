package utils

import (
	"errors"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

func ValidateExcel(file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	xlFile, err := excelize.OpenReader(f)
	if err != nil {
		return err
	}

	// Check if required columns exist
	headers, err := xlFile.GetRows("Sheet1")
	if err != nil {
		return err
	}
	if len(headers) < 1 || len(headers[0]) < 3 || headers[0][0] != "Name" || headers[0][1] != "Email" || headers[0][2] != "Phone" {
		return errors.New("invalid Excel format")
	}

	return nil
}
