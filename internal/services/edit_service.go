package services

import (
	"go-import-manage/internal/models"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
)

func EditService(id int, record models.Record) error {

	if err := mysql.UpdateRecord(id, record); err != nil {
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
