package events

import (
	"time"
)

type EventName string

const (
	UserRegisted EventName = "user.registered"
	UpdatedEmail EventName = "user.email.updated"
	SendSMS      EventName = "notifcation.sms.send"
)

type UpdatedEmailEvent struct {
	UserID   int
	NewEmail string // receiver's email
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
