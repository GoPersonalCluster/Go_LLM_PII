package models

import (
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID   string    `gorm:"size:64;not null"`
	OccuredAt int64     `gorm:"autoCreateTime"`
}

type TagExtraction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	PayloadID string    `gorm:"size:64;not null"`
	Result    string    `gorm:"size:64;not null"`
	OccuredAt int64     `gorm:"autoCreateTime"`
}
