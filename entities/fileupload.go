package entities

import (
	"time"

	"github.com/google/uuid"
)

type FileUpload struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"file_id"`
	Nama      string    `gorm:"not null" json:"nama"`
	Tipe      string    `json:"tipe,omitempty"`
	URL       string    `json:"url,omitempty"`
	KaryaID   uuid.UUID `gorm:"type:uuid;not null" json:"karya_id"`
	Karya     KaryaUKK  `gorm:"foreignKey:KaryaID" json:"karya,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
