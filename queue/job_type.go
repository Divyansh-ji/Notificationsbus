package queue

type NotificationJob struct {
	UserID  int
	Email   string
	Name    string
	Type    string // "email", "sms" etc
	Message string
}
