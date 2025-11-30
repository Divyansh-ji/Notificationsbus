package worker

import (
	"main.go/notifications"
	"main.go/queue"
)

func StartWorkers(n int) {
	for i := 0; i < n; i++ {
		go worker(i)
	}
}

func worker(id int) {
	for job := range queue.JobQueue {

		if job.Type == "email" {
			notifications.EmailService(job.UserID, job.Email, "Subject", "Body")
		}

		// later you can add SMS / webhook also
		// if job.Type == "sms" { services.SendSMS(...) }
	}
}
