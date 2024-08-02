package utils

import (
	"go-import-manage/internal/models"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

func ParseExcel(file *multipart.FileHeader) ([]models.Record, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	xlFile, err := excelize.OpenReader(f)
	if err != nil {
		return nil, err
	}

	records := []models.Record{}
	rows, err := xlFile.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] {
		record := models.Record{
			Name:  row[0],
			Email: row[1],
			Phone: row[2],
		}
		records = append(records, record)
	}
	return records, nil
}
