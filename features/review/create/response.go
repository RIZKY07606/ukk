package create

type CreateReviewResponse struct {
	ReviewID  string `json:"review_id"`
	Komentar  string `json:"komentar"`
	Rating    int    `json:"rating"`
	KaryaID   string `json:"karya_id"`
	CreatedAt string `json:"created_at"`
}

type CreateReviewResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    CreateReviewResponse `json:"data"`
}
