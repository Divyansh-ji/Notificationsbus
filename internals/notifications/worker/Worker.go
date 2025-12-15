package worker

import (
	"log"
	"sync"

	"main.go/events"
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

		switch job.Type {

		case "user_registered":
			go notifications.UserRegisterd(
				events.DefaultBus,
				job.UserID,
				job.Name,
				job.Email,
			)

		case "email & Sms":
			var wg sync.WaitGroup
			wg.Add(2)

			// EMAIL
			go func() {
				defer wg.Done()
				notifications.EmailService(
					job.UserID,
					job.Email,
					"Meesssage",
					"email is emailing ",
				)
			}()

			// SMS
			go func() {
				defer wg.Done()
				notifications.SmsServives(
					job.UserID,
					"-0987654321",
					"hey whatsapp",
				)
			}()

			wg.Wait()
		}
	}
}
func StartDeadLetterWorker() {
	go func() {
		for job := range queue.DeadLetterQueue {
			log.Println("Job failed permanently:", job)
		}
	}()
}
