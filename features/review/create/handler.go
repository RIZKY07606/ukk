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
//	@Summary		Buat review baru
//	@Description	Membuat data review untuk karya tertentu
//	@Tags			review
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Request body review"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/review [post]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Komentar == "" || req.Rating < 1 || req.Rating > 5 || req.KaryaID == "" || req.UserID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Field tidak valid atau kosong"})
		}

		karyaUUID, err := uuid.Parse(req.KaryaID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "KaryaID tidak valid"})
		}

		userUUID, err := uuid.Parse(req.UserID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "UserID tidak valid"})
		}

		review := entities.Review{
			ID:        uuid.New(),
			Komentar:  req.Komentar,
			Rating:    req.Rating,
			KaryaID:   karyaUUID,
			UserID:    userUUID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&review).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat review"})
		}

		return c.JSON(Response{
			ID:       review.ID,
			Komentar: review.Komentar,
			Rating:   review.Rating,
			KaryaID:  review.KaryaID,
			UserID:   review.UserID,
		})
	}
}
