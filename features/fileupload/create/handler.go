package create

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateFileUpload godoc
// @Summary     Create a new file upload
// @Description Create a new FileUpload untuk karya UKK
// @Tags        file_upload
// @Accept      json
// @Produce     json
// @Param       request body CreateFileUploadRequest true "FileUpload body"
// @Success     200 {object} CreateFileUploadResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/file-upload [post]
func CreateFileUpload(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateFileUploadRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if req.Nama == "" || req.URL == "" || req.KaryaID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Nama, URL, dan karya_id wajib diisi"})
		}
		fileID := uuid.New()
		karyaID, err := uuid.Parse(req.KaryaID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid karya_id UUID"})
		}
		f := entities.FileUpload{
			ID:        fileID,
			Nama:      req.Nama,
			Tipe:      req.Tipe,
			URL:       req.URL,
			KaryaID:   karyaID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Create(&f).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat file upload"})
		}
		res := CreateFileUploadResponse{
			FileID:    f.ID.String(),
			Nama:      f.Nama,
			Tipe:      f.Tipe,
			URL:       f.URL,
			KaryaID:   f.KaryaID.String(),
			CreatedAt: f.CreatedAt.Format(time.RFC3339),
		}
		return c.JSON(CreateFileUploadResponseWrapper{Code: 200, Message: "FileUpload berhasil dibuat", Data: res})
	}
}
