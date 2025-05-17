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
//
//	@Summary		Create Role
//	@Description	Create a new role
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request		true	"Role create request"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/role [post]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
		}

		if req.NIS == "" || req.Nama == "" || req.Kelas == "" || req.Jurusan == "" || req.UserID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Semua field wajib diisi"})
		}

		userUUID, err := uuid.Parse(req.UserID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "UserID tidak valid"})
		}

		siswa := entities.Siswa{
			ID:        uuid.New(),
			NIS:       req.NIS,
			Nama:      req.Nama,
			Kelas:     req.Kelas,
			Jurusan:   req.Jurusan,
			UserID:    userUUID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&siswa).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat siswa"})
		}

		return c.JSON(Response{
			ID:      siswa.ID,
			NIS:     siswa.NIS,
			Nama:    siswa.Nama,
			Kelas:   siswa.Kelas,
			Jurusan: siswa.Jurusan,
			UserID:  siswa.UserID,
		})
	}
}
