package notifications

import (
	"fmt"

	"time"

	"main.go/events"
)

func UserRegisterd(bus *events.EventBus, UserID int, Name string, email string) {
	info := events.UserRegisteredEvent{
		UserID:    UserID,
		Email:     email,
		Subject:   "User has been registerd",
		Name:      Name,
		CreatedAt: time.Now(),
	}
	fmt.Println("Info of the user registerd is", info)
	event := events.Event{
		Name:      "User.Registerd",
		Payload:   info,
		TimeStamp: time.Now(),
	}
	bus.Publish(event.Name, event)
}
func UserEmailUpdated(bus *events.Event, UserId int, NewEmail string) {
	updates := events.UpdatedEmailEvent{
		UserID:   UserId,
		NewEmail: NewEmail,
	}
	fmt.Println("Info of the User registerd is", updates)
	event := events.Event{
		Name:      "UserEmailUpdated",
		Payload:   updates,
		TimeStamp: time.Now(),
	}
	bus.Publish(event.Name, event)

}
