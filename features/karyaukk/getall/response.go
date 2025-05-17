package getall

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID         uuid.UUID `json:"karya_id"`
	Judul      string    `json:"judul"`
	Deskripsi  string    `json:"deskripsi,omitempty"`
	Link       string    `json:"link,omitempty"`
	SiswaID    uuid.UUID `json:"siswa_id"`
	KategoriID uuid.UUID `json:"kategori_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
