package notifications

import (
	"fmt"
	"time"
)

func EmailService(userID int, email string, subject string, message string) {
	fmt.Println("📧 Sending Email To:", email)
	fmt.Println("Subject:", subject)
	fmt.Println("Message:", message)
	fmt.Println("UserID:", userID)
	fmt.Println("Time:", time.Now())
}
func SmsServives(userID int, phone string, message string) {
	fmt.Println("Sending sms to:", userID)
	fmt.Println("phone:", phone)
	fmt.Println("message:", message)

}
