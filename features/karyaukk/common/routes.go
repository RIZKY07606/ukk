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
	groupKarya := api.Group("/karya-ukk")
	groupKarya.Post("/", create.Handler(db))
	groupKarya.Get("/", getall.Handler(db))
	groupKarya.Get("/:id", getbyid.Handler(db))
	groupKarya.Put("/:id", update.Handler(db))
	groupKarya.Delete("/:id", delete.Handler(db))

}
