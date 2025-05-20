package create

// CreateUserRequestBody
// @Description Create user request body
type CreateAdminRequest struct {
	// Your Name
	Nama string `json:"nama"`
	// Your Email
	Email string `json:"email" gorm:"unique"`
	// Your Password
	Password string `json:"password"`
}
