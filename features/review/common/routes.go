package common

import (
	"ukk-smkn2/features/review/create"
	"ukk-smkn2/features/review/delete"
	"ukk-smkn2/features/review/getall"
	"ukk-smkn2/features/review/getbyid"
	"ukk-smkn2/features/review/update"
	"ukk-smkn2/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("/review", middleware.JWTProtected())
	group.Post("/", create.CreateReview(db))
	group.Get("/", getall.GetAllReview(db))
	group.Get("/:id", getbyid.GetReviewByID(db))
	group.Put("/:id", update.UpdateReview(db))
	group.Delete("/:id", delete.DeleteReview(db))
}
