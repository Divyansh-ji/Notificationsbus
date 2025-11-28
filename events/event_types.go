package events

import (
	"time"
)

type SendEmailEvent struct {
	UserID    int       // optional, if you want to store notification later
	Email     string    // receiver's email
	Subject   string    // email subject
	Message   string    // body/content
	CreatedAt time.Time // timestamp when event was created
}
type SendSMSEvent struct {
	UserID    int
	Phone     string
	Message   string
	CreatedAt time.Time
}
type UserRegisteredEvent struct {
	UserID    int
	Email     string
	Subject   string
	Name      string
	CreatedAt time.Time
}
