package update

type UpdateReviewResponse struct {
	ID       string `json:"id"`
	Komentar string `json:"komentar"`
	Rating   int    `json:"rating"`
}

type UpdateReviewResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    UpdateReviewResponse `json:"data"`
}
