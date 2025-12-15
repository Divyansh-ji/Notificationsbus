package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/models"
	"main.go/queue"
)

func RegisterUser(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
		})
		return
	}
	if err := db.DB.Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
		})
		return
	}

	//	notifications.UserRegisterd(events.DefaultBus, User.ID, User.Name, User.Email) //this is the heavy task so Api become slow so we will make sure to add heavy task in jobqueue

	// job is created here
	// ⬇️ Create a job instead of doing heavy work here
	queue.JobQueue2 <- queue.Job{
		Type:   "user_registered",
		UserID: User.ID,
		Name:   User.Name,
		Email:  User.Email,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
	})

}
