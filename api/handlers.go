package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

func RegisterUser(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
		})
		return
	}
	
}
