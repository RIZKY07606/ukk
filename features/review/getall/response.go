package getall

type ReviewResponse struct {
	ReviewID  string `json:"review_id"`
	Komentar  string `json:"komentar"`
	Rating    int    `json:"rating"`
	KaryaID   string `json:"karya_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllReviewResponseWrapper struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []ReviewResponse `json:"data"`
}
