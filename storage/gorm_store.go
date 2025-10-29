package storage

import (
	"messages-api/models"

	"gorm.io/gorm"
)

type GormStore struct {
	DB *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	db.AutoMigrate(&models.Message{})
	return &GormStore{DB: db}
}

func (s *GormStore) Init() error {
	return nil
}

func (s *GormStore) Insert(msg models.Message) error {
	return s.DB.Create(&msg).Error
}

func (s *GormStore) GetAll() ([]models.Message, error) {
	var messages []models.Message
	err := s.DB.Find(&messages).Error
	return messages, err
}
