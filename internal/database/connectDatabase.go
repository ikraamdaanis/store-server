package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var err error

	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	fmt.Println("Successfully connected to the database.")
}
