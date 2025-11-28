package notifications

import "main/Notification/events"

func HandleUserRegistered(e events.Event) {
	data := e.Payload.(events.UserRegisteredEvent)

	// Now call the service
	EmailService(data.UserID, data.Email, data.Subject, data.Name)
}
