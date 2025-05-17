package getprofile

type GetUserResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email" gorm:"unique"`
}
