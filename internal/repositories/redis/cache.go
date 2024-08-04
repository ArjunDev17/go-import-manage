// package redis

// import (
// 	"encoding/json"
// 	"go-import-manage/internal/models"
// 	"log"
// )

// func CacheRecords(records []models.Record) error {
// 	data, err := json.Marshal(records)
// 	if err != nil {
// 		return err
// 	}

// 	err = RDB.Set(Ctx, "records", data, 0).Err()
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Records successfully cached in Redis")
// 	return nil
// }

// func GetCachedRecords() ([]models.Record, error) {
// 	data, err := RDB.Get(Ctx, "records").Result()
// 	if err != nil {
// 		log.Printf("Error retrieving records from Redis: %v", err)
// 		return nil, err
// 	}

// 	var records []models.Record
// 	if err := json.Unmarshal([]byte(data), &records); err != nil {
// 		log.Printf("Error unmarshalling records from Redis: %v", err)
// 		return nil, err
// 	}

//		log.Printf("Successfully retrieved %d records from Redis", len(records))
//		return records, nil
//	}
package redis

import (
	"encoding/json"
	"fmt"
	"go-import-manage/internal/models"
	"log"
)

func CacheRecords(records []models.Record) error {
	fmt.Println("Coming here ")
	data, err := json.Marshal(records)
	if err != nil {
		log.Printf("Error marshalling records: %v", err)
		return err
	}

	err = RDB.Set(Ctx, "records", data, 0).Err()
	if err != nil {
		log.Printf("Error caching records in Redis: %v", err)
		return err
	}

	log.Println("Records successfully cached in Redis")
	fmt.Println("------------------------------------------------------------------")
	return nil
}

func GetCachedRecords() ([]models.Record, error) {
	data, err := RDB.Get(Ctx, "records").Result()
	if err != nil {
		log.Printf("Error retrieving records from Redis: %v", err)
		return nil, err
	}

	var records []models.Record
	if err := json.Unmarshal([]byte(data), &records); err != nil {
		log.Printf("Error unmarshalling records from Redis: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved %d records from Redis", len(records))
	return records, nil
}
