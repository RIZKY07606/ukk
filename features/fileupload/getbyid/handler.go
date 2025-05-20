package getbyid

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetFileUploadByID godoc
// @Summary     Get file upload by ID
// @Description Ambil detail FileUpload berdasarkan ID
// @Tags        file_upload
// @Produce     json
// @Param       id path string true "FileUpload ID"
// @Success     200 {object} GetFileUploadByIDResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /api/file-upload/{id} [get]
func GetFileUploadByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}
		var f entities.FileUpload
		if err := db.First(&f, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "FileUpload tidak ditemukan"})
		}
		data := FileUploadDetailResponse{
			FileID:    f.ID.String(),
			Nama:      f.Nama,
			Tipe:      f.Tipe,
			URL:       f.URL,
			KaryaID:   f.KaryaID.String(),
			CreatedAt: f.CreatedAt.Format(time.RFC3339),
			UpdatedAt: f.UpdatedAt.Format(time.RFC3339),
		}
		return c.JSON(GetFileUploadByIDResponseWrapper{Code: 200, Message: "Success", Data: data})
	}
}
