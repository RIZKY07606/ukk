package create

import (
	"time"

	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler godoc
//
//	@Summary		Create Role
//	@Description	Create a new role
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Role create request"
//	@Success		201		{object}	Response
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/role [post]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		if req.Nama == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Nama role wajib diisi",
			})
		}

		role := entities.Role{
			ID:        uuid.New(),
			Nama:      req.Nama,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&role).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal membuat role",
			})
		}

		return c.JSON(Response{
			ID:   role.ID,
			Nama: role.Nama,
		})
	}
}
