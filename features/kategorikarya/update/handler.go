package update

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateKategoriKarya godoc
//
//	@Summary		Update kategori karya
//	@Description	Update kategori karya by ID
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Kategori ID"
//	@Param			request	body		UpdateKategoriKaryaRequest	true	"Update kategori karya body"
//	@Success		200		{object}	UpdateKategoriKaryaResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/kategori-karya/{id} [put]
func UpdateKategoriKarya(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid kategori ID"})
		}

		var req UpdateKategoriKaryaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.Nama == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Nama is required"})
		}

		var kategori entities.KategoriKarya
		if err := db.First(&kategori, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Kategori karya not found"})
		}

		kategori.Nama = req.Nama
		kategori.UpdatedAt = time.Now()

		if err := db.Save(&kategori).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update kategori karya"})
		}

		return c.JSON(UpdateKategoriKaryaResponseWrapper{
			Code:    200,
			Message: "Kategori karya updated successfully",
			Data: UpdateKategoriKaryaResponse{
				KategoriID: kategori.ID.String(),
				Nama:       kategori.Nama,
			},
		})
	}
}
