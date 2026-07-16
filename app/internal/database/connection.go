package database

import (
	"fmt"
	"log"

	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config := config.NewEnvironmentConfig()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
	)
	println("Connecting to database with connection string: ", connStr)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	err = DB.AutoMigrate(&models.Payload{}, &models.TagExtraction{})
	if err != nil {
		log.Fatal(err)
	}
}
