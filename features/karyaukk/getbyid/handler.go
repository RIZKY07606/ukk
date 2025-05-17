package getbyid

import (
	"net/http"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		karyaID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var karya entities.KaryaUKK
		if err := db.First(&karya, "id = ?", karyaID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Karya tidak ditemukan"})
		}

		return c.JSON(Response{
			ID:         karya.ID,
			Judul:      karya.Judul,
			Deskripsi:  karya.Deskripsi,
			Link:       karya.Link,
			SiswaID:    karya.SiswaID,
			KategoriID: karya.KategoriID,
			CreatedAt:  karya.CreatedAt,
			UpdatedAt:  karya.UpdatedAt,
		})
	}
}
