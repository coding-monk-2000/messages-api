package handlers

import (
	"bytes"
	"encoding/json"
	"messages-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMessages_Success(t *testing.T) {
	store := &MockStore{
		Messages: []models.Message{{Author: "Alice", Text: "Hello"}},
	}
	handler := &MessageHandler{Store: store}

	req := httptest.NewRequest(http.MethodGet, "/messages", nil)
	w := httptest.NewRecorder()

	handler.GetMessages(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestGetMessages_Failure(t *testing.T) {
	store := &MockStore{FailGet: true}
	handler := &MessageHandler{Store: store}

	req := httptest.NewRequest(http.MethodGet, "/messages", nil)
	w := httptest.NewRecorder()

	handler.GetMessages(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500, got %d", w.Code)
	}
}

func TestAddMessage_Success(t *testing.T) {
	store := &MockStore{}
	handler := &MessageHandler{Store: store}

	msg := models.Message{Author: "Bob", Text: "Hi"}
	body, _ := json.Marshal(msg)

	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.AddMessage(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %d", w.Code)
	}
}

func TestAddMessage_InvalidBody(t *testing.T) {
	store := &MockStore{}
	handler := &MessageHandler{Store: store}

	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader([]byte("not json")))
	w := httptest.NewRecorder()

	handler.AddMessage(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", w.Code)
	}
}

func TestAddMessage_MissingFields(t *testing.T) {
	store := &MockStore{}
	handler := &MessageHandler{Store: store}

	msg := models.Message{Author: "", Text: ""}
	body, _ := json.Marshal(msg)

	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.AddMessage(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", w.Code)
	}
}

func TestAddMessage_InsertFailure(t *testing.T) {
	store := &MockStore{FailInsert: true}
	handler := &MessageHandler{Store: store}

	msg := models.Message{Author: "Eve", Text: "Oops"}
	body, _ := json.Marshal(msg)

	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.AddMessage(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 Internal Server Error, got %d", w.Code)
	}
}
