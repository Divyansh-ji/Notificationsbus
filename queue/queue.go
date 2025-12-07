package queue

var JobQueue chan NotificationJob
var JobQueue2 chan Job

func InitQueue(size int) {
	JobQueue = make(chan NotificationJob, size)
	JobQueue2 = make(chan Job, size)
}
