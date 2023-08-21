package api

import (
	"log"
	"net/http"
	"os"
)

func RunServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Println("Server started on port " + port + ".")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
