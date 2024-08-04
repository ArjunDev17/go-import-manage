// package services

// import (
// 	"go-import-manage/internal/models"
// 	"go-import-manage/internal/repositories/mysql"
// 	"go-import-manage/internal/repositories/redis"
// 	"log"
// )

// func ViewService() ([]models.Record, error) {
// 	records, err := redis.GetCachedRecords()
// 	if err != nil || len(records) == 0 {
// 		log.Println("No records found in Redis, fetching from MySQL")
// 		records, err = mysql.GetRecords()
// 		if err != nil {
// 			return nil, err
// 		}
// 		log.Printf("Successfully retrieved %d records from MySQL", len(records))
// 	}

//		return records, nil
//	}
package services

import (
	"go-import-manage/internal/models"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
	"log"
)

func ViewService() ([]models.Record, error) {
	records, err := redis.GetCachedRecords()
	if err != nil || len(records) == 0 {
		log.Println("No records found in Redis, fetching from MySQL")
		records, err = mysql.GetRecords()
		if err != nil {
			return nil, err
		}
		log.Printf("Successfully retrieved %d records from MySQL", len(records))
	} else {
		log.Printf("Successfully retrieved %d records from Redis", len(records))
	}

	return records, nil
}
