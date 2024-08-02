package main

import (
	"go-import-manage/config"
	"go-import-manage/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize routes
	handlers.InitRoutes(r)

	// Load configuration
	config.LoadConfig()

	// Start the server
	if err := r.Run(":" + config.ServerPort); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
