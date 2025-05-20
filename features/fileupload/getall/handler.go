package getall

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllFileUpload godoc
// @Summary     Get all file uploads
// @Description Ambil semua FileUpload
// @Tags        file_upload
// @Produce     json
// @Success     200 {object} GetAllFileUploadResponseWrapper
// @Failure     500 {object} map[string]string
// @Router      /api/file-upload [get]
func GetAllFileUpload(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var list []entities.FileUpload
		if err := db.Find(&list).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data file upload"})
		}
		var resp []FileUploadResponse
		for _, f := range list {
			resp = append(resp, FileUploadResponse{
				FileID:    f.ID.String(),
				Nama:      f.Nama,
				Tipe:      f.Tipe,
				URL:       f.URL,
				KaryaID:   f.KaryaID.String(),
				CreatedAt: f.CreatedAt.Format(time.RFC3339),
				UpdatedAt: f.UpdatedAt.Format(time.RFC3339),
			})
		}
		return c.JSON(GetAllFileUploadResponseWrapper{Code: 200, Message: "Success", Data: resp})
	}
}
