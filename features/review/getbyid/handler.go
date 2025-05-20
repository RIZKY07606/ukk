package getbyid

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetReviewByID godoc
// @Summary     Get review by ID
// @Description Ambil detail review berdasarkan ID
// @Tags        review
// @Produce     json
// @Param       id path string true "Review ID"
// @Success     200 {object} GetReviewByIDResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /api/review/{id} [get]
func GetReviewByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
		}
		var rev entities.Review
		if err := db.First(&rev, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Review tidak ditemukan"})
		}
		data := ReviewDetailResponse{
			ReviewID:  rev.ID.String(),
			Komentar:  rev.Komentar,
			Rating:    rev.Rating,
			KaryaID:   rev.KaryaID.String(),
			CreatedAt: rev.CreatedAt.Format(time.RFC3339),
			UpdatedAt: rev.UpdatedAt.Format(time.RFC3339),
		}
		return c.JSON(GetReviewByIDResponseWrapper{Code: 200, Message: "Success", Data: data})
	}
}
