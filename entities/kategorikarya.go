package entities

import (
	"time"

	"github.com/google/uuid"
)

type KategoriKarya struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"kategori_id"`
	Nama      string     `gorm:"not null" json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	KaryaUKK  []KaryaUKK `gorm:"foreignKey:KategoriID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"karya_ukk,omitempty"`
}
