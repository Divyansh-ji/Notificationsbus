package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectTODB() {
	var err error

	dsn := os.Getenv("DB_URL")

	if dsn == " " {
		log.Fatal("the env variables are emptu hence we cannot connect to the database", err)
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error has been occured connectiong to the datbase", err)
	}

	log.Println("Successfully connected to the database")
}
