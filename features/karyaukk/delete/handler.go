package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteKaryaUKK godoc
//
//	@Summary		Delete karya UKK
//	@Description	Delete karya UKK by ID
//	@Tags			karya_ukk
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"KaryaUKK ID"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/karya-ukk/{id} [delete]
func DeleteKaryaUKK(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		if err := db.Delete(&entities.KaryaUKK{}, "id = ?", id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete karya UKK"})
		}

		return c.JSON(fiber.Map{
			"code":    200,
			"message": "Karya UKK deleted successfully",
		})
	}
}
