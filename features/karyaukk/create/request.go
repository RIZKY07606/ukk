package create

import "github.com/google/uuid"

type Request struct {
	Judul      string    `json:"judul"`
	Deskripsi  string    `json:"deskripsi,omitempty"`
	Link       string    `json:"link,omitempty"`
	SiswaID    uuid.UUID `json:"siswa_id"`
	KategoriID uuid.UUID `json:"kategori_id"`
}
