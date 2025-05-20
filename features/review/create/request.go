package create

type CreateReviewRequest struct {
	Komentar string `json:"komentar"`
	Rating   int    `json:"rating"`
	KaryaID  string `json:"karya_id"`
}
