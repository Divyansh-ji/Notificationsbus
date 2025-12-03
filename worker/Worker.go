package worker

import (
	"sync"

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

		if job.Type == "email & Sms" {
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				notifications.EmailService(job.UserID, job.Email, "Meesssage", "email is emailing ")
				wg.Done()
			}()

			go func() {
				notifications.SmsServives(job.UserID, "-0987654321", "hey whatsapp")
				wg.Done()
			}()

			wg.Wait()

		}
	}
}
