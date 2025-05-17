package create

import (
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
// @Summary      Create FileUpload
// @Description  Membuat data file upload baru
// @Tags         FileUpload
// @Accept       json
// @Produce      json
// @Param        data  body      Request  true  "Data file upload"
// @Success      200   {object}  Response
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /fileupload [post]

func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if req.Nama == "" || req.KaryaID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Nama dan KaryaID harus diisi"})
		}

		karyaUUID, err := uuid.Parse(req.KaryaID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "KaryaID tidak valid"})
		}

		fileUpload := entities.FileUpload{
			ID:        uuid.New(),
			Nama:      req.Nama,
			Tipe:      req.Tipe,
			URL:       req.URL,
			KaryaID:   karyaUUID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&fileUpload).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan file"})
		}

		return c.JSON(Response{
			ID:      fileUpload.ID,
			Nama:    fileUpload.Nama,
			Tipe:    fileUpload.Tipe,
			URL:     fileUpload.URL,
			KaryaID: fileUpload.KaryaID,
		})
	}
}
