package update

import (
	"net/http"
	"time"
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
//	@Summary		Update kategori karya
//	@Description	Mengubah nama kategori karya berdasarkan ID
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Kategori ID"
//	@Param			request	body		Request			true	"Request Body"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/kategori-karya/{id} [put]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		kategoriID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Nama == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Nama kategori wajib diisi"})
		}

		var kategori entities.KategoriKarya
		if err := db.First(&kategori, "id = ?", kategoriID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Kategori karya tidak ditemukan"})
		}

		kategori.Nama = req.Nama
		kategori.UpdatedAt = time.Now()

		if err := db.Save(&kategori).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengupdate kategori karya"})
		}

		return c.JSON(Response{
			ID:   kategori.ID,
			Nama: kategori.Nama,
		})
	}
}
