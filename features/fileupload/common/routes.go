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
	groupFileUpload := api.Group("/fileupload")
	groupFileUpload.Post("/", create.Handler(db))
	groupFileUpload.Get("/", getall.Handler(db))
	groupFileUpload.Get("/:id", getbyid.Handler(db))
	groupFileUpload.Put("/:id", update.Handler(db))
	groupFileUpload.Delete("/:id", delete.Handler(db))

}
