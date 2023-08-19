package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Println("Server started on port " + port + ".")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
