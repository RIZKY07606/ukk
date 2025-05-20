package update

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateSiswa godoc
//
//	@Summary		Update siswa
//	@Description	Update a siswa by ID
//	@Tags			siswa
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Siswa ID"
//	@Param			request	body		UpdateSiswaRequest	true	"Siswa update body"
//	@Success		200		{object}	UpdateSiswaResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/siswa/{id} [put]
func UpdateSiswa(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var req UpdateSiswaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.NIS == "" || req.Nama == "" || req.Kelas == "" || req.Jurusan == "" {
			return c.Status(400).JSON(fiber.Map{"error": "All fields are required"})
		}

		var siswa entities.Siswa
		if err := db.First(&siswa, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(404).JSON(fiber.Map{"error": "Siswa not found"})
			}
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch siswa"})
		}

		siswa.NIS = req.NIS
		siswa.Nama = req.Nama
		siswa.Kelas = req.Kelas
		siswa.Jurusan = req.Jurusan
		siswa.UpdatedAt = time.Now()

		if err := db.Save(&siswa).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update siswa"})
		}

		return c.JSON(UpdateSiswaResponseWrapper{
			Code:    200,
			Message: "Siswa updated successfully",
			Data: UpdateSiswaResponse{
				SiswaID: siswa.ID.String(),
				NIS:     siswa.NIS,
				Nama:    siswa.Nama,
				Kelas:   siswa.Kelas,
				Jurusan: siswa.Jurusan,
			},
		})
	}
}
