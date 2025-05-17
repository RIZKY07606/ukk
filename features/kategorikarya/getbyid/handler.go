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
//	@Summary		Ambil kategori karya berdasarkan ID
//	@Description	Mengembalikan detail kategori karya berdasarkan UUID
//	@Tags			kategori_karya
//	@Produce		json
//	@Param			id	path		string	true	"Kategori ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Router			/kategori-karya/{id} [get]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		kategoriID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var kategori entities.KategoriKarya
		if err := db.First(&kategori, "id = ?", kategoriID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Kategori karya tidak ditemukan"})
		}

		return c.JSON(Response{
			ID:   kategori.ID,
			Nama: kategori.Nama,
		})
	}
}
