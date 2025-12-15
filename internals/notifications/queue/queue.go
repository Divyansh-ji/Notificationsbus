package queue

var JobQueue chan NotificationJob
var JobQueue2 chan Job
var DeadLetterQueue = make(chan Job, 100)

func InitQueue(size int) {
	JobQueue = make(chan NotificationJob, size)
	JobQueue2 = make(chan Job, size)

}
