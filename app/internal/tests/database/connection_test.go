package database_test

import (
	"testing"
	"time"

	"github.com/GoPersonalCluster/go_llm_pii/app/internal/database"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/models"
)

func TestConnect(t *testing.T) {

	// Reset the global variable.
	database.DB = nil

	database.Connect()

	if database.DB == nil {
		t.Fatal("expected DB to be initialized")
	}

	sqlDB, err := database.DB.DB()
	time.Sleep(10 * time.Second) // Wait for the database to be ready
	if err != nil {
		t.Fatalf("failed to retrieve sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("database ping failed: %v", err)
	}
}

func TestCreatePayload(t *testing.T) {
	// Reset the global variable.
	database.DB = nil

	database.Connect()

	if database.DB == nil {
		t.Fatal("expected DB to be initialized")
	}

	sqlDB, err := database.DB.DB()
	if err != nil {
		t.Fatalf("failed to retrieve sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("database ping failed: %v", err)
	}
	entity := &models.Payload{
		EventID: "sdsd",
		Payload: "sdsd",
	}

	if err := database.DB.Create(entity).Error; err != nil {
		t.Fatalf("failed to insert payload: %v", err)
	}
}
