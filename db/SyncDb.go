package db

import "main.go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Notification{}, &models.User{})
}
