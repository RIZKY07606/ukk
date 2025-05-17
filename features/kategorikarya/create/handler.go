package create

import (
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
//	@Summary		Buat Kategori Karya
//	@Description	Endpoint untuk membuat kategori karya baru
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			body	body		Request	true	"Data kategori karya"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/kategori-karya [post]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Nama == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Nama kategori wajib diisi"})
		}

		kategori := entities.KategoriKarya{
			ID:        uuid.New(),
			Nama:      req.Nama,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&kategori).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat kategori karya"})
		}

		return c.JSON(Response{
			ID:   kategori.ID,
			Nama: kategori.Nama,
		})
	}
}
