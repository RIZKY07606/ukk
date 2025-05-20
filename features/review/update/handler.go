package update

import (
	"time"

	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateReview godoc
// @Summary     Update review
// @Description Update komentar & rating review
// @Tags        review
// @Accept      json
// @Produce     json
// @Param       id path string true "Review ID"
// @Param       request body UpdateReviewRequest true "Body update review"
// @Success     200 {object} UpdateReviewResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/review/{id} [put]
func UpdateReview(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}

		var req UpdateReviewRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.Komentar == "" || req.Rating < 1 || req.Rating > 5 {
			return c.Status(400).JSON(fiber.Map{"error": "Komentar & rating (1–5) wajib diisi"})
		}

		var rev entities.Review
		if err := db.First(&rev, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Review tidak ditemukan"})
		}

		rev.Komentar = req.Komentar
		rev.Rating = req.Rating
		rev.UpdatedAt = time.Now()

		if err := db.Save(&rev).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal update review"})
		}

		return c.JSON(UpdateReviewResponseWrapper{
			Code:    200,
			Message: "Review berhasil diupdate",
			Data: UpdateReviewResponse{
				ID:       rev.ID.String(),
				Komentar: rev.Komentar,
				Rating:   rev.Rating,
			},
		})
	}
}
