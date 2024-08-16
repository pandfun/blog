package main

import (
	"blog/database"
	"blog/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Invoked at start
func init() {

	database.ConnectDB()
}

func main() {

	sqlDB, err := database.DBConn.DB()
	if err != nil {
		log.Fatal("Error in sql connection")
	}

	defer sqlDB.Close()

	app := fiber.New()
	app.Use(logger.New())

	router.SetUpRoutes(app)

	log.Println("Server: Starting")
	if err := app.Listen(":9002"); err != nil {
		log.Fatal("Server: ", err.Error())
	}
}
