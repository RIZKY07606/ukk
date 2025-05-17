package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler godoc
//
//	@Summary		Ambil Semua Review
//	@Description	Mengambil semua review yang ada di database
//	@Tags			review
//	@Produce		json
//	@Success		200	{array}		Response
//	@Failure		500	{object}	ErrorResponse
//	@Router			/review [get]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reviews []entities.Review
		if err := db.Find(&reviews).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data review"})
		}

		var responses []Response
		for _, review := range reviews {
			responses = append(responses, Response{
				ID:       review.ID,
				Komentar: review.Komentar,
				Rating:   review.Rating,
				KaryaID:  review.KaryaID,
				UserID:   review.UserID,
			})
		}

		return c.JSON(responses)
	}
}
