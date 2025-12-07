package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/db"
	"main.go/events"
	"main.go/notifications"
	"main.go/queue"
	"main.go/worker"
)

func main() {
	db.LoadEnvVar()

	db.ConnectTODB()

	db.SyncDatabase()

	r := gin.Default()
	//Initialize queue and workers BEFORE starting the server
	queue.InitQueue(100)
	worker.StartWorkers(25)

	// Use the shared bus and subscribe BEFORE server starts
	bus := events.DefaultBus
	bus.Subscribe("User.Registerd", notifications.HandleUserRegistered)
	bus.Subscribe("UserEmailUpdated", notifications.HandleUserRegistered)

	r.POST("/User", api.RegisterUser)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)

		//notifications.UserRegisterd(bus, 101, "Divyansh", "divyanshTiwary01@gmail.com")

	}
}
