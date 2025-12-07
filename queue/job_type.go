package queue

type NotificationJob struct {
	UserID  int
	Email   string
	Name    string
	phone   string
	Type    string // "email", "sms" etc
	Message string
}

type Job struct {
	Type   string
	UserID int
	Name   string
	Email  string
}
