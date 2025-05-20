package getbyid

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetSiswaByID godoc
//
//	@Summary		Get siswa by ID
//	@Description	Get a siswa detail by ID
//	@Tags			siswa
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Siswa ID"
//	@Success		200		{object}	GetSiswaResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/siswa/{id} [get]
func GetSiswaByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var siswa entities.Siswa
		if err := db.First(&siswa, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(404).JSON(fiber.Map{"error": "Siswa not found"})
			}
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch siswa"})
		}

		return c.JSON(GetSiswaResponseWrapper{
			Code:    200,
			Message: "Success",
			Data: GetSiswaResponse{
				SiswaID: siswa.ID.String(),
				NIS:     siswa.NIS,
				Nama:    siswa.Nama,
				Kelas:   siswa.Kelas,
				Jurusan: siswa.Jurusan,
			},
		})
	}
}
