package getbyid

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetKaryaUKKByID godoc
//
//	@Summary		Get karya UKK by ID
//	@Description	Get detail karya UKK by ID
//	@Tags			karya_ukk
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"KaryaUKK ID"
//	@Success		200		{object}	GetKaryaUKKByIDResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Router			/api/karya-ukk/{id} [get]
func GetKaryaUKKByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		var karya entities.KaryaUKK
		if err := db.Preload("Siswa").Preload("Kategori").Preload("Admin").
			First(&karya, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Karya UKK not found"})
		}

		return c.JSON(GetKaryaUKKByIDResponseWrapper{
			Code:    200,
			Message: "Success",
			Data: GetKaryaUKKByIDResponse{
				KaryaID:    karya.ID.String(),
				Judul:      karya.Judul,
				Deskripsi:  karya.Deskripsi,
				Thumbnail:  karya.Thumbnail,
				Link:       karya.Link,
				SiswaID:    karya.SiswaID.String(),
				KategoriID: karya.KategoriID.String(),
				AdminID:    karya.AdminID.String(),
			},
		})
	}
}
