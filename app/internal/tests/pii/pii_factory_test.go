package pii

import (
	"testing"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/database"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/models"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/pii"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}

	err = db.AutoMigrate(&models.Payload{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	database.DB = db
}

func TestCreateStrategy_DefaultPiiEvent(t *testing.T) {
	setupDatabase(t)

	factory := &pii.PiiFactory{}

	event := &consumer.IntegrationEvent{
		ID:        "123",
		EventName: "DefaultPiiEvent",
		Payload:   "My name is John Michael Smith. I was born on March 14, 1992. My email address is john.smith@example.com and my personal phone number is +1 (555) 123-4567. My home address is 1234 Maple Avenue, Springfield, IL 62704. My Social Security Number is 123-45-6789, and my driver's license number is D12345678. My passport number is X1234567. My credit card number is 4111 1111 1111 1111 with expiration date 12/29 and CVV 123. My bank account number is 9876543210 and routing number is 021000021. Please send all correspondence to john.smith@example.com.",
		MetaHeader: []consumer.MetaHeader{
			{
				EventName: "DefaultPiiEvent",
			},
		},
	}

	handler, err := factory.CreateStrategy(event)
	println("handler: start")
	handler.Start()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if handler == nil {
		t.Fatal("expected strategy handler")
	}

	var count int64
	database.DB.Model(&models.Payload{}).Count(&count)

	if count != 1 {
		t.Fatalf("expected payload saved, got %d", count)
	}
}

func TestCreateStrategy_InvalidEvent(t *testing.T) {
	setupDatabase(t)

	factory := &pii.PiiFactory{}

	event := &consumer.IntegrationEvent{
		ID:        "123",
		EventName: "UnknownEvent",
		Payload:   "payload",
		MetaHeader: []consumer.MetaHeader{
			{
				EventName: "UnknownEvent",
			},
		},
	}

	handler, err := factory.CreateStrategy(event)

	if err == nil {
		t.Fatal("expected error")
	}

	if handler != nil {
		t.Fatal("expected nil strategy")
	}
}
