package common

import (
	"ukk-smkn2/features/siswa/create"
	"ukk-smkn2/features/siswa/delete"
	"ukk-smkn2/features/siswa/getall"
	"ukk-smkn2/features/siswa/getbyid"
	"ukk-smkn2/features/siswa/update"
	"ukk-smkn2/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("/siswa", middleware.JWTProtected())
	group.Post("/", create.CreateSiswa(db))
	group.Get("/", getall.GetAllSiswa(db))
	group.Get("/:id", getbyid.GetSiswaByID(db))
	group.Put("/:id", update.UpdateSiswa(db))
	group.Delete("/:id", delete.DeleteSiswa(db))
}
