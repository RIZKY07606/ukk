package create

type CreateUserResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email" gorm:"unique"`
}

type CreateUserResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateUserResponse `json:"data"`
}
