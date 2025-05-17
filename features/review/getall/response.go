package getall

import "github.com/google/uuid"

type Response struct {
	ID       uuid.UUID `json:"review_id"`
	Komentar string    `json:"komentar"`
	Rating   int       `json:"rating"`
	KaryaID  uuid.UUID `json:"karya_id"`
	UserID   uuid.UUID `json:"user_id"`
}
