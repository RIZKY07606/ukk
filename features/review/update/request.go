package update

type UpdateReviewRequest struct {
	Komentar string `json:"komentar"`
	Rating   int    `json:"rating"`
}
