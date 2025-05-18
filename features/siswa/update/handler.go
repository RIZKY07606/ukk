package update

import (
	"net/http"
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler mengupdate data siswa berdasarkan ID
//
// @Summary      Update Siswa by ID
// @Description  Memperbarui data siswa dengan ID tertentu
// @Tags         siswa
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Siswa ID"
// @Param        body  body      Request true  "Request body data siswa"
// @Success      200   {object}  Response
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/siswa/{id} [put]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		siswaID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Request body tidak valid"})
		}

		if req.NIS == "" || req.Nama == "" || req.Kelas == "" || req.Jurusan == "" || req.UserID == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Semua field wajib diisi"})
		}

		userUUID, err := uuid.Parse(req.UserID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "UserID tidak valid"})
		}

		var siswa entities.Siswa
		if err := db.First(&siswa, "id = ?", siswaID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Siswa tidak ditemukan"})
		}

		siswa.NIS = req.NIS
		siswa.Nama = req.Nama
		siswa.Kelas = req.Kelas
		siswa.Jurusan = req.Jurusan
		siswa.UserID = userUUID
		siswa.UpdatedAt = time.Now()

		if err := db.Save(&siswa).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengupdate siswa"})
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
