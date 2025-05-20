package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Siswa struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"siswa_id"`
	NIS       string         `gorm:"unique;not null" json:"nis"`
	Nama      string         `gorm:"not null" json:"nama"`
	Kelas     string         `gorm:"not null" json:"kelas"`
	Jurusan   string         `gorm:"not null" json:"jurusan"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	KaryaUKK  []KaryaUKK     `gorm:"foreignKey:SiswaID" json:"karya_ukk,omitempty"`
}
