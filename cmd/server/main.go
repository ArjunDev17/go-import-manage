package main

import (
	"go-import-manage/config"
	"go-import-manage/internal/handlers"
	"go-import-manage/internal/repositories/mysql"
	"go-import-manage/internal/repositories/redis"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load .env file manually
	err := godotenv.Load("/media/arjun/863684ab-ea66-44f7-9b95-f624c9361dea1/GoLang/projects/test_assignment/a1/go-import-manage/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	// Initialize routes
	handlers.InitRoutes(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Load configuration
	config.LoadConfig()

	// Initialize database connections
	mysql.InitMySQL()
	redis.InitRedis() // Ensure Redis is initialized

	// Start the server
	if err := r.Run(":" + config.ServerPort); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
