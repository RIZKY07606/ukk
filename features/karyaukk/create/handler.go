package create

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error" example:"Deskripsi error"`
}

// Handler godoc
// @Summary      Create a new KaryaUKK
// @Description  Membuat karya UKK baru dengan judul, deskripsi, link, siswa, dan kategori
// @Tags         KaryaUKK
// @Accept       json
// @Produce      json
// @Param        request  body      Request  true  "Request body untuk membuat karya UKK"
// @Success      200      {object}  Response
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /karyaukk [post]

func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Judul == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Judul wajib diisi"})
		}

		// Optional: kamu bisa tambah cek apakah SiswaID dan KategoriID ada di database

		karya := entities.KaryaUKK{
			ID:         uuid.New(),
			Judul:      req.Judul,
			Deskripsi:  req.Deskripsi,
			Link:       req.Link,
			SiswaID:    req.SiswaID,
			KategoriID: req.KategoriID,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := db.Create(&karya).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat karya"})
		}

		return c.JSON(Response{
			ID:         karya.ID,
			Judul:      karya.Judul,
			Deskripsi:  karya.Deskripsi,
			Link:       karya.Link,
			SiswaID:    karya.SiswaID,
			KategoriID: karya.KategoriID,
		})
	}
}
