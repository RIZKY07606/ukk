package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler retrieves all siswa records
//
// @Summary      Get All Siswa
// @Description  Mengambil semua data siswa
// @Tags         siswa
// @Produce      json
// @Success      200  {array}   Response
// @Failure      500  {object}  ErrorResponse
// @Router       /api/siswa [get]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var siswas []entities.Siswa
		if err := db.Find(&siswas).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data siswa"})
		}

		var responses []Response
		for _, siswa := range siswas {
			responses = append(responses, Response{
				ID:      siswa.ID,
				NIS:     siswa.NIS,
				Nama:    siswa.Nama,
				Kelas:   siswa.Kelas,
				Jurusan: siswa.Jurusan,
				UserID:  siswa.UserID,
			})
		}

		return c.JSON(responses)
	}
}
