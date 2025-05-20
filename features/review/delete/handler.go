package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteReview godoc
// @Summary     Delete review
// @Description Hapus review berdasarkan ID
// @Tags        review
// @Produce     json
// @Param       id path string true "Review ID"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/review/{id} [delete]
func DeleteReview(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		if _, err := uuid.Parse(idStr); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}
		if err := db.Delete(&entities.Review{}, "id = ?", idStr).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus review"})
		}
		return c.JSON(fiber.Map{"message": "Review berhasil dihapus"})
	}
}
