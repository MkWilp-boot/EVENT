package main

import (
	database "Event/cmd/pkg/database"
	structs "Event/cmd/pkg/structs"
	routes "Event/cmd/routes"

	"github.com/gofiber/fiber/v2"
)

const IsDevelopment = true

func main() {
	if IsDevelopment {
		database.OpenContext("test.db")
	}
	database.DB.AutoMigrate(&structs.Event{})

	// API
	app := fiber.New()
	api := app.Group("/api")
	{
		events := api.Group("/Events")
		{
			events.Get("/:id", routes.EventGetById)
			events.Post("/", routes.EventPost)
		}
	}
	app.Listen(":8080")
}
