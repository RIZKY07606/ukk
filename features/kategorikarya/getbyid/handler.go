package getbyid

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetKategoriKaryaByID godoc
//
//	@Summary		Get kategori karya by ID
//	@Description	Get kategori karya detail by ID
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Kategori ID"
//	@Success		200		{object}	GetKategoriKaryaByIDResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Router			/api/kategori-karya/{id} [get]
func GetKategoriKaryaByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid kategori ID"})
		}

		var kategori entities.KategoriKarya
		if err := db.First(&kategori, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Kategori karya not found"})
		}

		return c.JSON(GetKategoriKaryaByIDResponseWrapper{
			Code:    200,
			Message: "Success",
			Data: GetKategoriKaryaByIDResponse{
				KategoriID: kategori.ID.String(),
				Nama:       kategori.Nama,
			},
		})
	}
}
