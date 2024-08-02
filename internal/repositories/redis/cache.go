package redis

import (
	"encoding/json"
	"time"

	"go-import-manage/internal/models"
)

func CacheRecords(records []models.Record) error {
	data, err := json.Marshal(records)
	if err != nil {
		return err
	}

	return RDB.Set(Ctx, "records", data, 5*time.Minute).Err()
}

func GetCachedRecords() ([]models.Record, error) {
	data, err := RDB.Get(Ctx, "records").Bytes()
	if err != nil {
		return nil, err
	}

	var records []models.Record
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}
	return records, nil
}
