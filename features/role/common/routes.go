package common

import (
	"ukk-smkn2/features/role/create"
	"ukk-smkn2/features/role/delete"
	"ukk-smkn2/features/role/getall"
	"ukk-smkn2/features/role/getbyid"
	"ukk-smkn2/features/role/update"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("/role")

	group.Post("/", create.Handler(db))
	group.Get("/", getall.Handler(db))
	group.Get("/:id", getbyid.Handler(db))
	group.Put("/:id", update.Handler(db))
	group.Delete("/:id", delete.Handler(db))
}
