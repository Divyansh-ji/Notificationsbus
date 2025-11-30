package queue

var JobQueue chan NotificationJob

func InitQueue(size int) {
	JobQueue = make(chan NotificationJob, size)
}
