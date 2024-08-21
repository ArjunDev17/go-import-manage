package main

import (
	"go-import-manage/config"
	"go-import-manage/internal/handlers"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("/media/arjun/863684ab-ea66-44f7-9b95-f624c9361dea1/GoLang/go-Test/go-import-manage/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	r := gin.Default()

	handlers.InitRoutes(r)

	config.LoadConfig()

	mysql.InitMySQL()
	redis.InitRedis()

	if err := r.Run(":" + config.ServerPort); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
