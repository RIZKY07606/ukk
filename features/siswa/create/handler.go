package create

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateSiswa godoc
//
//	@Summary		Create siswa
//	@Description	Create a new siswa
//	@Tags			siswa
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateSiswaRequest	true	"Siswa body"
//	@Success		200		{object}	CreateSiswaResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/siswa [post]
func CreateSiswa(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateSiswaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.NIS == "" || req.Nama == "" || req.Kelas == "" || req.Jurusan == "" {
			return c.Status(400).JSON(fiber.Map{"error": "All fields are required"})
		}

		siswa := entities.Siswa{
			ID:        uuid.New(),
			NIS:       req.NIS,
			Nama:      req.Nama,
			Kelas:     req.Kelas,
			Jurusan:   req.Jurusan,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&siswa).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create siswa"})
		}

		return c.JSON(CreateSiswaResponseWrapper{
			Code:    200,
			Message: "Siswa created successfully",
			Data: CreateSiswaResponse{
				SiswaID: siswa.ID.String(),
				NIS:     siswa.NIS,
				Nama:    siswa.Nama,
				Kelas:   siswa.Kelas,
				Jurusan: siswa.Jurusan,
			},
		})
	}
}
