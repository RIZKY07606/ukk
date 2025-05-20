package create

type CreateAdminResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email" gorm:"unique"`
}

type CreateAdminResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateAdminResponse `json:"data"`
}
