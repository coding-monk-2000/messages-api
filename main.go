package main

import (
	"log"
	"net/http"

	"messages-api/handlers"
	"messages-api/storage"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	db, _ := storage.InitDatabase()
	r := mux.NewRouter()

	h := &handlers.MessageHandler{Store: db}

	r.HandleFunc("/messages", h.GetMessages).Methods("GET")
	r.HandleFunc("/messages", h.AddMessage).Methods("POST")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
