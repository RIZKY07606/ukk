package getbyid

import (
	"net/http"
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
//	@Summary		Get Role by ID
//	@Description	Get single role by ID
//	@Tags			role
//	@Produce		json
//	@Param			id	path		string	true	"Role ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Router			/api/role/{id} [get]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		roleID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var role entities.Role
		if err := db.First(&role, "id = ?", roleID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Role tidak ditemukan"})
		}

		return c.JSON(Response{
			ID:   role.ID,
			Nama: role.Nama,
		})
	}
}
