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

// @title Go Import Manage API
// @version 1.0
// @description This is a sample server for managing imported data.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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
