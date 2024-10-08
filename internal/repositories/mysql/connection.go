package mysql

import (
	"database/sql"
	"fmt"
	"go-import-manage/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
		config.MySQLDB,
	)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v\n", err)
	}

	// Set the maximum number of open connections
	DB.SetMaxOpenConns(100)
	// Set the maximum number of idle connections
	DB.SetMaxIdleConns(10)
	// Set the maximum lifetime of a connection
	DB.SetConnMaxLifetime(0)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Could not ping MySQL: %v\n", err)
	}

	log.Println("Connected to MySQL successfully")
}
