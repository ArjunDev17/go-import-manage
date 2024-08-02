package services

import (
	"go-import-manage/internal/models"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
)

func ViewService() ([]models.Record, error) {
	records, err := redis.GetCachedRecords()
	if err == nil && len(records) > 0 {
		return records, nil
	}

	records, err = mysql.GetRecords()
	if err != nil {
		return nil, err
	}

	if err := redis.CacheRecords(records); err != nil {
		return nil, err
	}

	return records, nil
}
