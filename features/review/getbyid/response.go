package getbyid

type GetReviewByIDResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    ReviewDetailResponse `json:"data"`
}

type ReviewDetailResponse struct {
	ReviewID  string `json:"review_id"`
	Komentar  string `json:"komentar"`
	Rating    int    `json:"rating"`
	KaryaID   string `json:"karya_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
