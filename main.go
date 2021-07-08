package main

import (
	"log"
	"os"

	"github.com/ahmedkrdzalic/StackAndFlow/database"
	"github.com/ahmedkrdzalic/StackAndFlow/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Printf("PORT is: " + port)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":" + port)
}
