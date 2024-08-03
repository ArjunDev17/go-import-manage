package redis

import (
	"context"
	"fmt"
	"go-import-manage/config"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	// Convert RedisDB from string to int
	dbNum, err := strconv.Atoi(config.RedisDB)
	if err != nil {
		log.Fatalf("Invalid RedisDB value: %v", err)
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword, // No password set
		DB:       dbNum,                // Use the converted integer value
	})
	_, err = RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")
}
