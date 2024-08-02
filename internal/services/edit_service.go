package services

import (
	"go-import-manage/internal/models"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
	"log"
	"strconv"
)

func EditService(id string, record models.Record) error {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		return err
	}
	if err := mysql.UpdateRecord(intID, record); err != nil {
		return err
	}

	records, err := mysql.GetRecords()
	if err != nil {
		return err
	}

	if err := redis.CacheRecords(records); err != nil {
		return err
	}

	return nil
}
