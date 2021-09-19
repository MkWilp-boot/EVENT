package main

import (
	database "Event/cmd/pkg/database"
	structs "Event/cmd/pkg/structs"
	"fmt"
	"time"
)

const IsDevelopment = true

func main() {
	if IsDevelopment {
		database.OpenContext("test.db")
	}
	database.DB.AutoMigrate(&structs.Event{})
	database.DB.Create(&structs.Event{
		Name:     "Evento Teste",
		Image:    "evt_" + time.Now().GoString(),
		When:     time.Now().Add(time.Duration(time.Now().Day() + 30)),
		Where:    "Rio de Janeiro",
		Duration: time.Duration(time.Hour * 2),
	})
	var e structs.Event
	database.DB.First(&e, 1)
	fmt.Println(e)
}
