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

// Handler mengupdate data siswa berdasarkan ID
//
// @Summary      Update Siswa by ID
// @Description  Memperbarui data siswa dengan ID tertentu
// @Tags         siswa
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Siswa ID"
// @Param        body  body      Request true  "Request body data siswa"
// @Success      200   {object}  Response
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /siswa/{id} [put]
// @Security     BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		roleID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if req.Nama == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Nama role wajib diisi"})
		}

		var role entities.Role
		if err := db.First(&role, "id = ?", roleID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Role tidak ditemukan"})
		}

		role.Nama = req.Nama
		role.UpdatedAt = time.Now()

		if err := db.Save(&role).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengupdate role"})
		}

		return c.JSON(Response{
			ID:   role.ID,
			Nama: role.Nama,
		})
	}
}
