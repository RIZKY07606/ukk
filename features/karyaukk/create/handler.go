package create

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateKaryaUKK godoc
//
//	@Summary		Create karya UKK
//	@Description	Create a new karya UKK
//	@Tags			karya_ukk
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateKaryaUKKRequest	true	"KaryaUKK body"
//	@Success		200		{object}	CreateKaryaUKKResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/karya-ukk [post]
func CreateKaryaUKK(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateKaryaUKKRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		// validasi minimal: Judul, SiswaID, KategoriID, AdminID
		if req.Judul == "" || req.SiswaID == "" || req.KategoriID == "" || req.AdminID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Judul, SiswaID, KategoriID, dan AdminID wajib diisi"})
		}

		idSiswa, err1 := uuid.Parse(req.SiswaID)
		idKategori, err2 := uuid.Parse(req.KategoriID)
		idAdmin, err3 := uuid.Parse(req.AdminID)
		if err1 != nil || err2 != nil || err3 != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format UUID tidak valid"})
		}

		karya := entities.KaryaUKK{
			ID:         uuid.New(),
			Judul:      req.Judul,
			Deskripsi:  req.Deskripsi,
			Thumbnail:  req.Thumbnail,
			Link:       req.Link,
			SiswaID:    idSiswa,
			KategoriID: idKategori,
			AdminID:    idAdmin,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := db.Create(&karya).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat karya UKK"})
		}

		return c.JSON(CreateKaryaUKKResponseWrapper{
			Code:    200,
			Message: "Karya UKK created successfully",
			Data: CreateKaryaUKKResponse{
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
