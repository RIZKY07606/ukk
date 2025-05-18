package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// Handler godoc
// @Summary      Delete FileUpload by ID
// @Description  Menghapus data FileUpload berdasarkan ID
// @Tags         FileUpload
// @Produce      json
// @Param        id   path      string  true  "FileUpload ID (UUID)"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/fileupload/{id} [delete]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ID harus diisi"})
		}

		res := db.Delete(&entities.FileUpload{}, "id = ?", id)
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus file"})
		}

		if res.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "File tidak ditemukan"})
		}

		return c.JSON(fiber.Map{"message": "File berhasil dihapus"})
	}
}
