package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteKategoriKarya godoc
//
//	@Summary		Delete kategori karya
//	@Description	Delete kategori karya by ID
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Kategori ID"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/kategori-karya/{id} [delete]
func DeleteKategoriKarya(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid kategori ID"})
		}

		if err := db.Delete(&entities.KategoriKarya{}, "id = ?", id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete kategori karya"})
		}

		return c.JSON(fiber.Map{"message": "Kategori karya deleted successfully"})
	}
}
