package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteFileUpload godoc
// @Summary     Delete file upload
// @Description Hapus FileUpload berdasarkan ID
// @Tags        file_upload
// @Produce     json
// @Param       id path string true "FileUpload ID"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/file-upload/{id} [delete]
func DeleteFileUpload(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		if _, err := uuid.Parse(idStr); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}
		if err := db.Delete(&entities.FileUpload{}, "id = ?", idStr).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus file upload"})
		}
		return c.JSON(fiber.Map{"message": "FileUpload berhasil dihapus"})
	}
}
