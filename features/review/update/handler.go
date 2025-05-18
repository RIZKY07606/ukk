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

// Handler godoc
//
//	@Summary		Update Review by ID
//	@Description	Mengupdate data review berdasarkan ID
//	@Tags			review
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Review ID"
//	@Param			body	body		Request		true	"Data review yang akan diupdate"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/review/{id} [put]
//	@Security		BearerAuth
func Handler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		reviewID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Request body invalid"})
		}

		if req.Komentar == "" || req.Rating < 1 || req.Rating > 5 || req.KaryaID == "" || req.UserID == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Field tidak valid atau kosong"})
		}

		karyaUUID, err := uuid.Parse(req.KaryaID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "KaryaID tidak valid"})
		}

		userUUID, err := uuid.Parse(req.UserID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "UserID tidak valid"})
		}

		var review entities.Review
		if err := db.First(&review, "id = ?", reviewID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Review tidak ditemukan"})
		}

		review.Komentar = req.Komentar
		review.Rating = req.Rating
		review.KaryaID = karyaUUID
		review.UserID = userUUID
		review.UpdatedAt = time.Now()

		if err := db.Save(&review).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengupdate review"})
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
