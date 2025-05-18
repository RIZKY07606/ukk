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
//
//	@Summary		Hapus Kategori Karya
//	@Description	Hapus satu kategori karya berdasarkan ID
//	@Tags			kategori_karya
//	@Param			id	path		string	true	"ID Kategori Karya"
//	@Produce		json
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/kategori-karya/{id} [delete]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ID harus diisi"})
		}

		res := db.Delete(&entities.KategoriKarya{}, "id = ?", id)
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus kategori karya"})
		}

		if res.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Kategori karya tidak ditemukan"})
		}

		return c.JSON(fiber.Map{"message": "Kategori karya berhasil dihapus"})
	}
}
