package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handler godoc
// @Summary      Get all KaryaUKK
// @Description  Mengambil semua data karya UKK
// @Tags         KaryaUKK
// @Produce      json
// @Success      200  {array}   Response
// @Failure      500  {object}  map[string]string{"error": "Gagal mengambil data karya"}
// @Router       /karyaukk [get]

func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var karya []entities.KaryaUKK
		if err := db.Find(&karya).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data karya"})
		}

		var responses []Response
		for _, k := range karya {
			responses = append(responses, Response{
				ID:         k.ID,
				Judul:      k.Judul,
				Deskripsi:  k.Deskripsi,
				Link:       k.Link,
				SiswaID:    k.SiswaID,
				KategoriID: k.KategoriID,
				CreatedAt:  k.CreatedAt,
				UpdatedAt:  k.UpdatedAt,
			})
		}

		return c.JSON(responses)
	}
}
