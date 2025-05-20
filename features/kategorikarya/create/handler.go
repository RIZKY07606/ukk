package create

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateKategoriKarya godoc
//
//	@Summary		Create kategori karya
//	@Description	Create a new kategori karya
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateKategoriKaryaRequest	true	"Kategori Karya body"
//	@Success		200		{object}	CreateKategoriKaryaResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/kategori-karya [post]
func CreateKategoriKarya(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateKategoriKaryaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.Nama == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Nama is required"})
		}

		kategori := entities.KategoriKarya{
			ID:        uuid.New(),
			Nama:      req.Nama,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&kategori).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create kategori karya"})
		}

		return c.JSON(CreateKategoriKaryaResponseWrapper{
			Code:    200,
			Message: "Kategori karya created successfully",
			Data: CreateKategoriKaryaResponse{
				KategoriID: kategori.ID.String(),
				Nama:       kategori.Nama,
			},
		})
	}
}
