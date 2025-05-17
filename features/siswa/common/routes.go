package common

import (
	"ukk-smkn2/features/siswa/create"
	"ukk-smkn2/features/siswa/delete"
	"ukk-smkn2/features/siswa/getall"
	"ukk-smkn2/features/siswa/getbyid"
	"ukk-smkn2/features/siswa/update"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("/siswa")

	group.Post("/", create.Handler(db))
	group.Get("/", getall.Handler(db))
	group.Get("/:id", getbyid.Handler(db))
	group.Put("/:id", update.Handler(db))
	group.Delete("/:id", delete.Handler(db))
}
