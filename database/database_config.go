package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() (*sql.DB, error) {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading env file")
		}
	}

	// Retrieve database credentials from environment variables
	DBUser := os.Getenv("DBUser")
	DBPass := os.Getenv("DBPass")
	DBHost := os.Getenv("DBHost")
	DBPort := os.Getenv("DBPort")
	DBName := os.Getenv("DBName")

	// Create the data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser, DBPass, DBHost, DBPort, DBName)

	// Open the database connection using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}

	return sqlDb, err
}
