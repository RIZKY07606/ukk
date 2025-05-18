package getbyid

import (
	"net/http"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler retrieves siswa by ID
//
// @Summary      Get Siswa by ID
// @Description  Mengambil data siswa berdasarkan ID
// @Tags         siswa
// @Produce      json
// @Param        id   path      string  true  "Siswa ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/siswa/{id} [get]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		siswaID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var siswa entities.Siswa
		if err := db.First(&siswa, "id = ?", siswaID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Siswa tidak ditemukan"})
		}

		return c.JSON(Response{
			ID:      siswa.ID,
			NIS:     siswa.NIS,
			Nama:    siswa.Nama,
			Kelas:   siswa.Kelas,
			Jurusan: siswa.Jurusan,
			UserID:  siswa.UserID,
		})
	}
}
