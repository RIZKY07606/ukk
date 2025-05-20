package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"review_id"`
	Komentar  string         `gorm:"type:text" json:"komentar"`
	Rating    int            `gorm:"not null" json:"rating"`
	KaryaID   uuid.UUID      `gorm:"type:uuid;not null" json:"karya_id"`
	Karya     *KaryaUKK      `gorm:"foreignKey:KaryaID" json:"karya,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
