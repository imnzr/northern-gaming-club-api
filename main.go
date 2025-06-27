package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imnzr/northern-gaming-club-api/database"
)

func main() {
	// env := godotenv.Load()
	// if env != nil {
	// 	fmt.Println("env file is not loaded")
	// }

	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("connection to database failed")
	}

	defer db.Close()

	fmt.Println("connection to database successfully")

}
