package delete

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
//	@Summary		Hapus Review
//	@Description	Hapus review berdasarkan ID
//	@Tags			review
//	@Param			id	path		string	true	"Review ID"
//	@Success		200	{object}	map[string]string
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/review/{id} [delete]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ID harus diisi"})
		}

		res := db.Delete(&entities.Review{}, "id = ?", id)
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus review"})
		}

		if res.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Review tidak ditemukan"})
		}

		return c.JSON(fiber.Map{"message": "Review berhasil dihapus"})
	}
}
