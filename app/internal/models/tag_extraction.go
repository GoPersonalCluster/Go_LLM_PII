package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payload struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID   string    `gorm:"size:64;not null"`
	OccuredAt int64     `gorm:"autoCreateTime"`
	Payload   []byte    `gorm:"type:string;not null"`

	// Relacionamento 1:N
	TagExtractions []TagExtraction `gorm:"foreignKey:PayloadID;references:ID"`
}

func (p *Payload) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	return nil
}

type TagExtraction struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Chave estrangeira
	PayloadID uuid.UUID `gorm:"type:uuid;not null;index"`

	Result    string `gorm:"size:64;not null"`
	OccuredAt int64  `gorm:"autoCreateTime"`

	// Relacionamento N:1
	Payload Payload `gorm:"foreignKey:PayloadID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *TagExtraction) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
