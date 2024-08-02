package services

import (
	"fmt"
	"go-import-manage/internal/models"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
	"go-import-manage/internal/utils"
	"mime/multipart"
	"sync"
)

func ImportService(file *multipart.FileHeader) error {
	records, err := utils.ParseExcel(file)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, record := range records {
		wg.Add(1)
		go func(record models.Record) {
			defer wg.Done()
			if err := mysql.InsertRecord(record); err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
		}(record)
	}

	wg.Wait()

	if len(errs) > 0 {
		return fmt.Errorf("errors occurred during import: %v", errs)
	}

	if err := redis.CacheRecords(records); err != nil {
		return err
	}

	return nil
}
