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

// Handler godoc
// @Summary      Update a KaryaUKK by ID
// @Description  Mengupdate karya UKK berdasarkan ID yang diberikan
// @Tags         KaryaUKK
// @Accept       json
// @Produce      json
// @Param        id     path      string   true  "KaryaUKK ID (UUID)"
// @Param        data   body      Request  true  "Data karya UKK yang akan diupdate"
// @Success      200    {object}  Response
// @Failure      400    {object}  ErrorResponse
// @Failure      404    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /api/karyaukk/{id} [put]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		karyaID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Judul == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Judul wajib diisi"})
		}

		// Parse SiswaID dan KategoriID ke uuid.UUID
		siswaID, err := uuid.Parse(req.SiswaID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "SiswaID tidak valid"})
		}

		kategoriID, err := uuid.Parse(req.KategoriID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "KategoriID tidak valid"})
		}

		var karya entities.KaryaUKK
		if err := db.First(&karya, "id = ?", karyaID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Karya tidak ditemukan"})
		}

		karya.Judul = req.Judul
		karya.Deskripsi = req.Deskripsi
		karya.Link = req.Link
		karya.SiswaID = siswaID
		karya.KategoriID = kategoriID
		karya.UpdatedAt = time.Now()

		if err := db.Save(&karya).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal update karya"})
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
