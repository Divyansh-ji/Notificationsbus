package main

import (
	"main.go/events"
	"main.go/notifications"
)

func main() {
	bus := events.NewEventBus()

	bus.Subscribe("User.Registerd", notifications.HandleUserRegistered)

	notifications.UserRegisterd(bus, 101, "Divyansh", "divyanshTiwary01@gmail.com")

}
