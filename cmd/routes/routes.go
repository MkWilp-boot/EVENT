package routes

import (
	"Event/cmd/pkg/database"
	"Event/cmd/pkg/structs"
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func EventGetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(nil)
	}
	var event structs.Event
	database.DB.First(&event, id)
	return c.JSON(event)
}

func EventPost(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var payload structs.Event
	reader := bytes.NewReader(c.Body())
	decoder := json.NewDecoder(reader)

	decoder.Decode(&payload)

	database.DB.Create(&payload).First(&payload)
	log.Println(payload)
	return c.JSON(payload)
}
