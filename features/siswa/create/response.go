package create

import (
	"github.com/google/uuid"
)

type Response struct {
	ID      uuid.UUID `json:"siswa_id"`
	NIS     string    `json:"nis"`
	Nama    string    `json:"nama"`
	Kelas   string    `json:"kelas"`
	Jurusan string    `json:"jurusan"`
	UserID  uuid.UUID `json:"user_id"`
}
