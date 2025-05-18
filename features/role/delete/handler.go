package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// Handler godoc
//
//	@Summary		Delete Role
//	@Description	Delete a role by ID
//	@Tags			role
//	@Produce		json
//	@Param			id	path		string	true	"Role ID (UUID)"
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/role/{id} [delete]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		// validasi UUID
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID tidak valid (harus UUID)",
			})
		}

		res := db.Delete(&entities.Role{}, "id = ?", id)
		if res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menghapus role",
			})
		}
		if res.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Role tidak ditemukan",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Role berhasil dihapus",
		})
	}
}
