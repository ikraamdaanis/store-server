package main

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/ikraamdaanis/store-server/internal/database"
	"github.com/ikraamdaanis/store-server/internal/initializers"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadVariables()
	database.ConnectDatabase()
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string
	Email   string
	Session string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Migrate the schema
	err := database.DB.AutoMigrate(&Product{}, &User{})

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
