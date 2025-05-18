package delete

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler deletes siswa by ID
//
// @Summary      Delete Siswa
// @Description  Menghapus data siswa berdasarkan ID
// @Tags         siswa
// @Param        id   path      string  true  "Siswa ID"
// @Produce      json
// @Success      200  {object}  MessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/siswa/{id} [delete]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ID tidak boleh kosong"})
		}

		res := db.Delete(&entities.Siswa{}, "id = ?", id)
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus siswa"})
		}

		if res.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Siswa tidak ditemukan"})
		}

		return c.JSON(fiber.Map{"message": "Siswa berhasil dihapus"})
	}
}
