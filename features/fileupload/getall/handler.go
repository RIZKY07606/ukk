package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler godoc
// @Summary      Get All FileUploads
// @Description  Mengambil semua data FileUpload
// @Tags         FileUpload
// @Produce      json
// @Success      200  {array}   Response
// @Failure      500  {object}  ErrorResponse
// @Router       /api/fileupload [get]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var files []entities.FileUpload
		if err := db.Find(&files).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data file"})
		}

		var responses []Response
		for _, file := range files {
			responses = append(responses, Response{
				ID:      file.ID,
				Nama:    file.Nama,
				Tipe:    file.Tipe,
				URL:     file.URL,
				KaryaID: file.KaryaID,
			})
		}

		return c.JSON(responses)
	}
}
