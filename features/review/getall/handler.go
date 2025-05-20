package getall

import (
	"time"
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllReview godoc
// @Summary     Get all reviews
// @Description Ambil semua review
// @Tags        review
// @Produce     json
// @Success     200 {object} GetAllReviewResponseWrapper
// @Failure     500 {object} map[string]string
// @Router      /api/review [get]
func GetAllReview(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var list []entities.Review
		if err := db.Find(&list).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data review"})
		}
		var resp []ReviewResponse
		for _, rev := range list {
			resp = append(resp, ReviewResponse{
				ReviewID:  rev.ID.String(),
				Komentar:  rev.Komentar,
				Rating:    rev.Rating,
				KaryaID:   rev.KaryaID.String(),
				CreatedAt: rev.CreatedAt.Format(time.RFC3339),
				UpdatedAt: rev.UpdatedAt.Format(time.RFC3339),
			})
		}
		return c.JSON(GetAllReviewResponseWrapper{Code: 200, Message: "Success", Data: resp})
	}
}
