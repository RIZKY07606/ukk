package getbyid

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
//
//	@Summary		Ambil Review by ID
//	@Description	Mengambil satu data review berdasarkan ID
//	@Tags			review
//	@Produce		json
//	@Param			id	path		string	true	"Review ID"
//	@Success		200	{object}	Response
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Router			/api/review/{id} [get]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		reviewID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var review entities.Review
		if err := db.First(&review, "id = ?", reviewID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Review tidak ditemukan"})
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
