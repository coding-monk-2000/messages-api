package handlers

import (
	"errors"
	"messages-api/models"
)

type MockStore struct {
	Messages   []models.Message
	FailGet    bool
	FailInsert bool
}

func (m *MockStore) Init() error {
	return nil
}

func (m *MockStore) GetAll() ([]models.Message, error) {
	if m.FailGet {
		return nil, errors.New("get failed")
	}
	return m.Messages, nil
}

func (m *MockStore) Insert(msg models.Message) error {
	if m.FailInsert {
		return errors.New("insert failed")
	}
	m.Messages = append(m.Messages, msg)
	return nil
}
