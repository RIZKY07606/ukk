package common

import (
	createKarya "ukk-smkn2/features/karyaukk/create"
	deleteKarya "ukk-smkn2/features/karyaukk/delete"
	getallKarya "ukk-smkn2/features/karyaukk/getall"
	getbyidKarya "ukk-smkn2/features/karyaukk/getbyid"
	updateKarya "ukk-smkn2/features/karyaukk/update"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	groupKarya := api.Group("/karya-ukk")
	groupKarya.Post("/", createKarya.Handler(db))
	groupKarya.Get("/", getallKarya.Handler(db))
	groupKarya.Get("/:id", getbyidKarya.Handler(db))
	groupKarya.Put("/:id", updateKarya.Handler(db))
	groupKarya.Delete("/:id", deleteKarya.Handler(db))
}
