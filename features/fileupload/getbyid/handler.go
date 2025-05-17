package getbyid

import (
	"net/http"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Handler godoc
// @Summary      Get FileUpload by ID
// @Description  Mengambil data FileUpload berdasarkan ID
// @Tags         FileUpload
// @Produce      json
// @Param        id   path      string  true  "FileUpload ID (UUID)"
// @Success      200  {object}  Response
// @Failure      400  {object}  map[string]string{"error": "ID tidak valid"}
// @Failure      404  {object}  map[string]string{"error": "File tidak ditemukan"}
// @Router       /fileupload/{id} [get]

func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		fileID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var file entities.FileUpload
		if err := db.First(&file, "id = ?", fileID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "File tidak ditemukan"})
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
