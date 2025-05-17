package getall

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
//	@Summary		Ambil semua kategori karya
//	@Description	Mengembalikan list semua kategori karya
//	@Tags			kategori_karya
//	@Produce		json
//	@Success		200	{array}		Response
//	@Failure		500	{object}	ErrorResponse
//	@Router			/kategori-karya [get]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var kategori []entities.KategoriKarya
		if err := db.Find(&kategori).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data kategori karya"})
		}

		var responses []Response
		for _, k := range kategori {
			responses = append(responses, Response{
				ID:   k.ID,
				Nama: k.Nama,
			})
		}

		return c.JSON(responses)
	}
}
