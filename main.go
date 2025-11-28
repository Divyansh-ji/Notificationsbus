package main

import (
	"main/Notification/events"
	"main/Notification/notifications"
)

func main() {
	bus := events.NewEventBus()

	bus.Subscribe("User.Registerd", notifications.HandleUserRegistered)

	notifications.UserRegisterd(bus, 101, "Divyansh", "divyanshTiwary01@gmail.com")

}
