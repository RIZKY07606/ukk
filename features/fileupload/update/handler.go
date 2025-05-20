package update

import (
	"time"

	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateFileUpload godoc
// @Summary     Update file upload
// @Description Update nama, tipe, dan URL FileUpload
// @Tags        file_upload
// @Accept      json
// @Produce     json
// @Param       id path string true "FileUpload ID"
// @Param       request body UpdateFileUploadRequest true "Update body"
// @Success     200 {object} UpdateFileUploadResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/file-upload/{id} [put]
func UpdateFileUpload(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		var req UpdateFileUploadRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if req.Nama == "" || req.URL == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Nama dan URL wajib diisi"})
		}

		var f entities.FileUpload
		if err := db.First(&f, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "FileUpload tidak ditemukan"})
		}

		f.Nama = req.Nama
		f.Tipe = req.Tipe
		f.URL = req.URL
		f.UpdatedAt = time.Now()

		if err := db.Save(&f).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal update file upload"})
		}

		// Build response data
		data := UpdateFileUploadResponse{
			FileID:    f.ID.String(),
			Nama:      f.Nama,
			Tipe:      f.Tipe,
			URL:       f.URL,
			KaryaID:   f.KaryaID.String(),
			UpdatedAt: f.UpdatedAt.Format(time.RFC3339),
		}

		return c.JSON(UpdateFileUploadResponseWrapper{
			Code:    200,
			Message: "FileUpload berhasil diupdate",
			Data:    data,
		})
	}
}
