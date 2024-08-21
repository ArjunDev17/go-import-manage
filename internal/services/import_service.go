package services

import (
	"fmt"
	"go-import-manage/internal/models"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
	"go-import-manage/internal/utils"
	"log"
	"mime/multipart"
	"sync"
)

func ImportService(file *multipart.FileHeader) error {
	// Open the file
	f, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Parse the file
	records, err := utils.ParseExcel(f)
	if err != nil {
		return fmt.Errorf("failed to parse Excel file: %w", err)
	}

	log.Printf("Parsed %d records from Excel file", len(records))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	sem := make(chan struct{}, 10) // limit to 10 concurrent goroutines

	// [GIN] 2024/08/08 - 17:41:35 | 200 |  665.957506ms |             ::1 | POST     "/api/import"  10
	// [GIN] 2024/08/08 - 17:42:30 | 200 |  617.228821ms |             ::1 | POST     "/api/import"  20
	// [GIN] 2024/08/08 - 17:43:06 | 200 |  417.088889ms |             ::1 | POST     "/api/import"  40
	// [GIN] 2024/08/08 - 17:45:04 | 200 |  357.246648ms |             ::1 | POST     "/api/import"  80

	for _, record := range records {
		wg.Add(1)
		go func(record models.Record) {
			defer wg.Done()
			sem <- struct{}{}        // acquire a token
			defer func() { <-sem }() // release the token

			if err := mysql.InsertRecord(record); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to insert record: %w", err))
				mu.Unlock()
			}
		}(record)
	}

	wg.Wait()

	if len(errs) > 0 {
		for _, err := range errs {
			log.Println(err)
		}
		return fmt.Errorf("errors occurred during import")
	}

	log.Println("All records inserted into MySQL")

	if err := redis.CacheRecords(records); err != nil {
		log.Printf("Failed to cache records: %v", err)
		return fmt.Errorf("failed to cache records: %w", err)
	}

	log.Printf("Successfully cached %d records in Redis", len(records))
	return nil
}
