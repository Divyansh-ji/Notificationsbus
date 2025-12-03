package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/events"
	"main.go/models"
	"main.go/notifications"
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

	notifications.UserRegisterd(events.DefaultBus, User.ID, User.Name, User.Email)
	
	// job is created here

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
	})

}
