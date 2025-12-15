package notifications

import (
	"fmt"
	"time"

	"main.go/queue"
)

func EmailService(userID int, email string, subject string, message string) error {
	fmt.Println("📧 Sending Email To:", email)
	fmt.Println("Subject:", subject)
	fmt.Println("Message:", message)
	fmt.Println("UserID:", userID)
	fmt.Println("Time:", time.Now())
	return nil
}
func SmsServives(userID int, phone string, message string) {
	fmt.Println("Sending sms to:", userID)
	fmt.Println("phone:", phone)
	fmt.Println("message:", message)

}
func RetryEmail(job queue.Job, maxRetries int) {
	for i := 0; i < maxRetries; i++ {
		err := EmailService(
			job.UserID,
			job.Email,
			"emails",
			"hey newUSer",
		)
		if err == nil {
			return //
		}
		//bacloff measn wqe will hold for the seconds so this will help not on spamming the request
		time.Sleep(time.Second * time.Duration(i+1))

	}
	queue.DeadLetterQueue <- job

}
