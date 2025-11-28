package models

import "time"

type Notification struct {
	ID        int
	UserID    int
	Message   string
	Type      string
	CreatedAt time.Time
}
