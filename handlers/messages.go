package handlers

import (
	"encoding/json"
	"messages-api/models"
	"messages-api/storage"
	"net/http"
)

type MessageHandler struct {
	Store storage.MessageStore
}

func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Store.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

func (h *MessageHandler) AddMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if msg.Author == "" || msg.Text == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if err := h.Store.Insert(msg); err != nil {
		http.Error(w, "Failed to insert message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}
