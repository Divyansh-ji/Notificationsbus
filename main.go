package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/db"
	"main.go/events"
	"main.go/notifications"
)

func main() {
	db.LoadEnvVar()

	db.ConnectTODB()

	db.SyncDatabase()

	bus := events.NewEventBus()

	bus.Subscribe("User.Registerd", notifications.HandleUserRegistered)

	notifications.UserRegisterd(bus, 101, "Divyansh", "divyanshTiwary01@gmail.com")

	r := gin.Default()

	r.POST("/User", api.RegisterUser)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
