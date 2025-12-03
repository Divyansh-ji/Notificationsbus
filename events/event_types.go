package events

import (
	"time"
)

type UpdatedEmailEvent struct {
	UserID   int    // optional, if you want to store notification later
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
