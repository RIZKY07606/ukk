package common

import (
	"ukk-smkn2/features/review/create"
	"ukk-smkn2/features/review/delete"
	"ukk-smkn2/features/review/getall"
	"ukk-smkn2/features/review/getbyid"
	"ukk-smkn2/features/review/update"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("/review")
	group.Post("/", create.Handler(db))
	group.Get("/", getall.Handler(db))
	group.Get("/:id", getbyid.Handler(db))
	group.Put("/:id", update.Handler(db))
	group.Delete("/:id", delete.Handler(db))
}
