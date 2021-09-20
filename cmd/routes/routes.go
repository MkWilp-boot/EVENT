package routes

import (
	"Event/cmd/pkg/database"
	"Event/cmd/pkg/structs"
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func EventGetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(404)
		return c.JSON(nil)
	}
	var event structs.Event
	dbErr := database.DB.First(&event, id)
	if dbErr.Error == gorm.ErrRecordNotFound {
		c.Status(404)
		return c.JSON(nil)
	}
	return c.JSON(event)
}

func EventPost(c *fiber.Ctx) error {
	c.Accepts("application/json")
	payload := deserializePayLoad(c)
	database.DB.Create(&payload).First(&payload)
	c.Status(201)
	return c.JSON(payload)
}

func EventDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(nil)
	}
	database.DB.Delete(&structs.Event{}, id)
	c.Status(204)
	return c.JSON(nil)
}

func deserializePayLoad(c *fiber.Ctx) structs.Event {
	var payload structs.Event
	reader := bytes.NewReader(c.Body())
	decoder := json.NewDecoder(reader)

	decoder.Decode(&payload)
	return payload
}
