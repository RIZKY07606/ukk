package update

import (
	"time"

	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateKaryaUKK godoc
//
//	@Summary		Update karya UKK
//	@Description	Update karya UKK by ID
//	@Tags			karya_ukk
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"KaryaUKK ID"
//	@Param			request	body		UpdateKaryaUKKRequest	true	"Update body"
//	@Success		200		{object}	UpdateKaryaUKKResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/karya-ukk/{id} [put]
func UpdateKaryaUKK(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		var req UpdateKaryaUKKRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		var karya entities.KaryaUKK
		if err := db.First(&karya, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Karya UKK not found"})
		}

		karya.Judul = req.Judul
		karya.Deskripsi = req.Deskripsi
		karya.Thumbnail = req.Thumbnail
		karya.Link = req.Link
		karya.UpdatedAt = time.Now()

		if err := db.Save(&karya).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update karya UKK"})
		}

		return c.JSON(UpdateKaryaUKKResponseWrapper{
			Code:    200,
			Message: "Karya UKK updated successfully",
			Data: UpdateKaryaUKKResponse{
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
