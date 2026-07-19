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
		Payload:   []byte("payload"),
		MetaHeader: []consumer.MetaHeader{
			{
				EventName: "DefaultPiiEvent",
			},
		},
	}

	handler, err := factory.CreateStrategy(event)

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
		Payload:   []byte("payload"),
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

func TestGetDefaultErrorResponse(t *testing.T) {
	factory := &pii.PiiFactory{}

	event := &consumer.IntegrationEvent{
		EventName: "UnknownEvent",
		MetaHeader: []consumer.MetaHeader{
			{
				EventName: "UnknownEvent",
			},
		},
	}

	err := factory.GetDefaultErrorResponse(event)

	if err == nil {
		t.Fatal("expected error")
	}

	expected := "UnknownEventevent not found"

	if err.Error() != expected {
		t.Fatalf("expected %q got %q", expected, err.Error())
	}
}

func TestGetDefaultPiiStrategy(t *testing.T) {
	factory := &pii.PiiFactory{}

	event := &consumer.IntegrationEvent{
		EventName: "DefaultPiiEvent",
		MetaHeader: []consumer.MetaHeader{
			{
				EventName: "DefaultPiiEvent",
			},
		},
	}

	handler, err := factory.GetDefaultPiiStrategy(event)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if handler == nil {
		t.Fatal("expected strategy")
	}
}
