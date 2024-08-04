// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// var (
// 	MySQLUser     string
// 	MySQLPassword string
// 	MySQLDB       string
// 	MySQLHost     string
// 	MySQLPort     string
// 	RedisHost     string
// 	RedisPort     string
// 	RedisPassword string
// 	RedisDB       string
// 	ServerPort    string
// )

// func LoadConfig() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found")
// 	}

// 	MySQLUser = getEnv("MYSQL_USER", "root")
// 	MySQLPassword = getEnv("MYSQL_PASSWORD", "password")
// 	MySQLDB = getEnv("MYSQL_DB", "your_database")
// 	MySQLHost = getEnv("MYSQL_HOST", "localhost")
// 	MySQLPort = getEnv("MYSQL_PORT", "3306")
// 	RedisHost = getEnv("REDIS_HOST", "localhost")
// 	RedisPort = getEnv("REDIS_PORT", "6379")
// 	RedisPassword = getEnv("REDIS_PASSWORD", "")
// 	RedisDB = getEnv("REDIS_DB", "0")
// 	ServerPort = getEnv("SERVER_PORT", "8080")

// 	log.Printf("MySQLUser: %s, MySQLPassword: %s, MySQLDB: %s, MySQLHost: %s, MySQLPort: %s",
// 		MySQLUser, MySQLPassword, MySQLDB, MySQLHost, MySQLPort)
// }

//	func getEnv(key, defaultValue string) string {
//		if value, exists := os.LookupEnv(key); exists {
//			return value
//		}
//		return defaultValue
//	}
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MySQLUser     string
	MySQLPassword string
	MySQLDB       string
	MySQLHost     string
	MySQLPort     string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string // Keep RedisDB as a string for now
	ServerPort    string
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	MySQLUser = getEnv("MYSQL_USER", "root")
	MySQLPassword = getEnv("MYSQL_PASSWORD", "password")
	MySQLDB = getEnv("MYSQL_DB", "your_database")
	MySQLHost = getEnv("MYSQL_HOST", "localhost")
	MySQLPort = getEnv("MYSQL_PORT", "3306")
	RedisHost = getEnv("REDIS_HOST", "localhost")
	RedisPort = getEnv("REDIS_PORT", "6379")
	RedisPassword = getEnv("REDIS_PASSWORD", "")
	RedisDB = getEnv("REDIS_DB", "0")
	ServerPort = getEnv("SERVER_PORT", "8080")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
