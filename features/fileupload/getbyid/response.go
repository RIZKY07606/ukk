package getbyid

import "github.com/google/uuid"

type Response struct {
	ID      uuid.UUID `json:"file_id"`
	Nama    string    `json:"nama"`
	Tipe    string    `json:"tipe,omitempty"`
	URL     string    `json:"url,omitempty"`
	KaryaID uuid.UUID `json:"karya_id"`
}
