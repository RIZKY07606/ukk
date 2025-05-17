package getbyid

import "github.com/google/uuid"

type Response struct {
	ID   uuid.UUID `json:"kategori_id"`
	Nama string    `json:"nama"`
}
