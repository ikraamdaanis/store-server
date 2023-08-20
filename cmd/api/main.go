package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ikraamdaanis/store-server/internal/database"
	"github.com/ikraamdaanis/store-server/internal/initializers"
)

func init() {
	initializers.LoadVariables()
	database.ConnectDatabase()
}

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Sessions  []Session `gorm:"foreignKey:AccountID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID        uint      `gorm:"primaryKey"`
	AccountID uuid.UUID `gorm:"type:uuid;not null"`
	Token     string    `gorm:"type:varchar(255);not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Migrate the schema
	err := database.DB.AutoMigrate(&Account{}, &Session{})

	if err != nil {
		log.Println("Error migrating database: ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Println("Server started on port " + port + ".")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
