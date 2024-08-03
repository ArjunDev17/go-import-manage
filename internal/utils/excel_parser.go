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
	// Open the Excel file using excelize
	xlFile, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}

	// Get the list of sheets
	sheetNames := xlFile.GetSheetMap()
	if len(sheetNames) == 0 {
		return nil, errors.New("no sheets found in the excel file")
	}

	// Attempt to use "Sheet1", otherwise use the first available sheet
	var sheetName string
	for _, name := range sheetNames {
		sheetName = name
		if name == "Sheet1" {
			break
		}
	}
	log.Printf("Using sheet: %s", sheetName)

	// Get the rows from the selected sheet
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet %s: %w", sheetName, err)
	}

	// Parse the rows into records
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

	// Check if records were parsed successfully
	if len(records) == 0 {
		return nil, errors.New("no valid records found")
	}

	return records, nil
}
