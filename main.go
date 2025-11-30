package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/db"
	"main.go/events"
	"main.go/notifications"
	"main.go/worker"
)

func main() {
	db.LoadEnvVar()

	db.ConnectTODB()

	db.SyncDatabase()

	r := gin.Default()

	r.POST("/User", api.RegisterUser)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	bus := events.NewEventBus()

	worker.StartWorkers(5)

	bus.Subscribe("User.Registerd", notifications.HandleUserRegistered)

	notifications.UserRegisterd(bus, 101, "Divyansh", "divyanshTiwary01@gmail.com")

}
