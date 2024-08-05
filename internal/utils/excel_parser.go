package utils

import (
	"errors"
	"fmt"
	"go-import-manage/internal/models"
	"io"
	"log"

	"github.com/xuri/excelize/v2"
)

// ParseExcel parses an Excel file and returns a slice of records
func ParseExcel(file io.Reader) ([]models.Record, error) {
	xlFile, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}

	sheetNames := xlFile.GetSheetMap()
	if len(sheetNames) == 0 {
		return nil, errors.New("no sheets found in the excel file")
	}

	var sheetName string
	for _, name := range sheetNames {
		sheetName = name
		if name == "Sheet1" {
			break
		}
	}
	log.Printf("Using sheet: %s", sheetName)

	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet %s: %w", sheetName, err)
	}

	var records []models.Record
	for i, row := range rows[1:] { // Skip the header row
		if len(row) < 10 {
			log.Printf("Skipping row %d: unexpected row length %d", i+2, len(row))
			log.Printf("Row content: %v", row)
			continue
		}

		record := models.Record{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			County:      row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		}
		records = append(records, record)
	}

	if len(records) == 0 {
		return nil, errors.New("no valid records found")
	}

	log.Printf("Parsed %d records from Excel file", len(records))
	return records, nil
}
