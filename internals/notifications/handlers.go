package notifications

import (
	"main.go/events"
	"main.go/queue"
)

func HandleUserRegistered(e events.Event) {

	data := e.Payload.(events.UserRegisteredEvent)

	queue.JobQueue <- queue.NotificationJob{
		UserID:  data.UserID,
		Email:   data.Email,
		Name:    data.Name,
		Type:    "email & Sms",
		Message: "Welcome to our platform 🎉",
	}

}
