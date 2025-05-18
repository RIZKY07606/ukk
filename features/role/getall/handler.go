package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler godoc
//
//	@Summary		Get All Roles
//	@Description	Get list of all roles
//	@Tags			role
//	@Produce		json
//	@Success		200	{array}	Response
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/role [get]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var roles []entities.Role
		if err := db.Find(&roles).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal mengambil data roles",
			})
		}

		var responses []Response
		for _, role := range roles {
			responses = append(responses, Response{
				ID:   role.ID,
				Nama: role.Nama,
			})
		}

		return c.JSON(responses)
	}
}
