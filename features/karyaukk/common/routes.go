package common

import (
	createKarya "ukk-smkn2/features/karyaukk/create"
	deleteKarya "ukk-smkn2/features/karyaukk/delete"
	getallKarya "ukk-smkn2/features/karyaukk/getall"
	getbyidKarya "ukk-smkn2/features/karyaukk/getbyid"
	updateKarya "ukk-smkn2/features/karyaukk/update"
	"ukk-smkn2/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	groupKarya := api.Group("/karya-ukk", middleware.JWTProtected())
	groupKarya.Post("/", createKarya.CreateKaryaUKK(db))
	groupKarya.Get("/", getallKarya.GetAllKaryaUKK(db))
	groupKarya.Get("/:id", getbyidKarya.GetKaryaUKKByID(db))
	groupKarya.Put("/:id", updateKarya.UpdateKaryaUKK(db))
	groupKarya.Delete("/:id", deleteKarya.DeleteKaryaUKK(db))
}
