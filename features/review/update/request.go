package update

type Request struct {
	Komentar string `json:"komentar"`
	Rating   int    `json:"rating"`
	KaryaID  string `json:"karya_id"`
	UserID   string `json:"user_id"`
}
