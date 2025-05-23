package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KaryaUKK struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"karya_id"`
	Judul      string         `gorm:"not null" json:"judul"`
	Deskripsi  string         `json:"deskripsi,omitempty"`
	Thumbnail  string         `json:"thumbnail,omitempty"` // path atau URL gambar
	Link       string         `json:"link,omitempty"`      // link ke repo / karya
	SiswaID    uuid.UUID      `gorm:"type:uuid;not null" json:"siswa_id"`
	Siswa      Siswa          `gorm:"foreignKey:SiswaID" json:"siswa,omitempty"`
	KategoriID uuid.UUID      `gorm:"type:uuid;not null" json:"kategori_id"`
	Kategori   KategoriKarya  `gorm:"foreignKey:KategoriID" json:"kategori,omitempty"`
	AdminID    uuid.UUID      `gorm:"type:uuid;not null" json:"admin_id"`
	Admin      Admin          `gorm:"foreignKey:AdminID" json:"admin,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	FileUpload []FileUpload   `gorm:"foreignKey:KaryaID" json:"file_upload,omitempty"`
	Review     []Review       `gorm:"foreignKey:KaryaID" json:"review,omitempty"`
}
