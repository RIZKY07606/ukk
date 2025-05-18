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
// @Summary      Update FileUpload by ID
// @Description  Mengupdate data FileUpload berdasarkan ID
// @Tags         FileUpload
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "FileUpload ID (UUID)"
// @Param        body  body      Request true  "Data file yang diupdate"
// @Success      200   {object}  Response
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/fileupload/{id} [put]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		fileID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Nama == "" || req.KaryaID == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Nama dan KaryaID harus diisi"})
		}

		karyaUUID, err := uuid.Parse(req.KaryaID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "KaryaID tidak valid"})
		}

		var file entities.FileUpload
		if err := db.First(&file, "id = ?", fileID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "File tidak ditemukan"})
		}

		file.Nama = req.Nama
		file.Tipe = req.Tipe
		file.URL = req.URL
		file.KaryaID = karyaUUID
		file.UpdatedAt = time.Now()

		if err := db.Save(&file).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengupdate file"})
		}

		return c.JSON(Response{
			ID:      file.ID,
			Nama:    file.Nama,
			Tipe:    file.Tipe,
			URL:     file.URL,
			KaryaID: file.KaryaID,
		})
	}
}
