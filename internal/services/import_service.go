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

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, record := range records {

		wg.Add(1)
		go func(record models.Record) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					mu.Lock()
					errs = append(errs, fmt.Errorf("panic occurred: %v", r))
					mu.Unlock()
				}
			}()
			if err := mysql.InsertRecord(record); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to insert record: %w", err))
				mu.Unlock()
			}
		}(record)
	}

	wg.Wait()
	fmt.Println("step 5")
	if len(errs) > 0 {

		for _, err := range errs {
			log.Println(err)
		}
		return fmt.Errorf("errors occurred during import")
	}

	fmt.Println("calling cache REcords")
	err = redis.CacheRecords(records)
	if err != nil {
		fmt.Printf("failed to cache records: %v\n", err)
		return fmt.Errorf("failed to cache records: %w", err)
	}
	fmt.Println("after calling cache REcords")

	log.Printf("Successfully cached %d records in Redis", len(records))
	return nil
}
