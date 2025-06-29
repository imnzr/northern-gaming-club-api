package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/imnzr/northern-gaming-club-api/database"
)

func main() {

	// Database Configuration START
	db, errDb := database.DatabaseConnection()
	if errDb != nil {
		fmt.Println("connection to database failed")
	}
	defer db.Close()
	fmt.Println("connection to database successfully")
	// Database Configuration FINISH

	// HTTP SERVER
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})
	// TEST HELLO WORLD
	app.Get("/", func(controller *fiber.Ctx) error {
		return controller.SendString("Hello World")
	})

	err := app.Listen("localhost:8080")
	if err != nil {
		panic(err)
	}
}
