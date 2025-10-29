package storage

import "messages-api/models"

type MessageStore interface {
	Init() error
	Insert(msg models.Message) error
	GetAll() ([]models.Message, error)
}
