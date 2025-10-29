// storage/gorm_store_test.go
package storage

import (
	"messages-api/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open in-memory DB: %v", err)
	}
	return db
}

func TestNewGormStore_AutoMigrate(t *testing.T) {
	db := setupTestDB(t)
	store := NewGormStore(db)

	// Check if the table exists
	if !db.Migrator().HasTable(&models.Message{}) {
		t.Error("Expected Message table to be auto-migrated")
	}
	if store == nil {
		t.Error("Expected store to be initialized")
	}
}

func TestInsertAndGetAll(t *testing.T) {
	db := setupTestDB(t)
	store := NewGormStore(db)

	msg := models.Message{Author: "Alice", Text: "Hello"}
	if err := store.Insert(msg); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	messages, err := store.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}
	if len(messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(messages))
	}
	if messages[0].Author != "Alice" || messages[0].Text != "Hello" {
		t.Errorf("Unexpected message content: %+v", messages[0])
	}
}

func TestInit(t *testing.T) {
	db := setupTestDB(t)
	store := NewGormStore(db)

	if err := store.Init(); err != nil {
		t.Errorf("Init should return nil, got %v", err)
	}
}
