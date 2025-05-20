package create

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateReview godoc
// @Summary     Create a new review
// @Description Create a new review untuk karya UKK
// @Tags        review
// @Accept      json
// @Produce     json
// @Param       request body CreateReviewRequest true "Review body"
// @Success     200 {object} CreateReviewResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/review [post]
func CreateReview(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateReviewRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		if req.Komentar == "" || req.Rating < 1 || req.Rating > 5 || req.KaryaID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Komentar, rating (1–5), dan karya_id wajib diisi"})
		}

		rev := entities.Review{
			ID:        uuid.New(),
			Komentar:  req.Komentar,
			Rating:    req.Rating,
			KaryaID:   mustParseUUID(req.KaryaID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Create(&rev).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat review"})
		}

		res := CreateReviewResponse{
			ReviewID:  rev.ID.String(),
			Komentar:  rev.Komentar,
			Rating:    rev.Rating,
			KaryaID:   rev.KaryaID.String(),
			CreatedAt: rev.CreatedAt.Format(time.RFC3339),
		}
		return c.JSON(CreateReviewResponseWrapper{Code: 200, Message: "Review berhasil dibuat", Data: res})
	}
}

// helper untuk parsing UUID
func mustParseUUID(s string) uuid.UUID {
	id, _ := uuid.Parse(s)
	return id
}
