package delete

import (
	"net/http"

	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler godoc
// @Summary      Delete a KaryaUKK by ID
// @Description  Menghapus karya UKK berdasarkan ID yang diberikan
// @Tags         KaryaUKK
// @Produce      json
// @Param        id   path      string  true  "KaryaUKK ID (UUID)"
// @Success      200  {object}  MessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /karyaukk/{id} [delete]
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		karyaID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		if err := db.Delete(&entities.KaryaUKK{}, "id = ?", karyaID).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus karya"})
		}

		return c.JSON(fiber.Map{"message": "Karya berhasil dihapus"})
	}
}
